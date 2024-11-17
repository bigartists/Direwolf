package maas

import (
	"context"
	"fmt"
	"github.com/bigartists/Direwolf/internal/module/conversation"
	"github.com/bigartists/Direwolf/internal/module/message"
	"github.com/bigartists/Direwolf/pkg/llm"
	"github.com/bigartists/Direwolf/pkg/result"
	"github.com/bigartists/Direwolf/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type MaasBaseInfo struct {
	Id       int64     `json:"id"`
	Model    string    `json:"model"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_time"`
}

type IMaasService interface {
	GetMaasList(userId int64) ([]*MaasBaseInfo, error)
	GetMaasDetail(id int64) *result.ErrorResult
	ImportMaas(model *MaasRequest, userId int64) (*Maas, error)
	UpdateMaas(id int, model *Maas) *result.ErrorResult
	DeleteMaas(id int) *result.ErrorResult
	InvokeMaas(ctx context.Context, req LLMRequest) (<-chan string, error)
	Completion(ctx *gin.Context, req CompletionResponse, userId int64)
}

type ModelService struct {
	repo                IModelRepo
	messageService      message.IMessageService
	conversationService conversation.IConversationService
}

func (m *ModelService) GetMaasList(userId int64) ([]*MaasBaseInfo, error) {
	modelList, err := m.repo.FindMaasList(userId)
	if err != nil {
		return nil, err
	}
	var modelBaseInfoList []*MaasBaseInfo
	for _, maas := range modelList {
		modelBaseInfoList = append(modelBaseInfoList, &MaasBaseInfo{
			Id:       maas.ID,
			Model:    maas.Model,
			Name:     maas.Name,
			Avatar:   maas.Avatar,
			CreateAt: maas.CreateTime,
		})
	}
	return modelBaseInfoList, nil
	//return modelList, err
}

func (m *ModelService) GetMaasDetail(id int64) *result.ErrorResult {
	model := &Maas{}
	model, err := m.repo.FindMaasById(id, model)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(model, nil)
}

func (m *ModelService) ImportMaas(modelRequest *MaasRequest, userId int64) (*Maas, error) {
	isDuplicate, err := m.repo.CheckDuplicateMaasBaseURL(context.Background(), modelRequest.Model, modelRequest.BaseURL)

	fmt.Println("isDuplicate:= ", isDuplicate)
	fmt.Println("err:= ", err)

	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, ErrDuplicateMaasBaseURL
	}
	fmt.Println(".................=======")
	model := NewMaas(
		WithModel(modelRequest.Model),
		WithApiKey(modelRequest.APIKey),
		WithBaseUrl(modelRequest.BaseURL),
		WithModelType(modelRequest.ModelType),
		WithName(modelRequest.Name),
		WithCreateBy(userId),
		WithAvatar(modelRequest.Avatar),
	)

	err = m.repo.ImportMaas(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *ModelService) UpdateMaas(id int, model *Maas) *result.ErrorResult {
	err := m.repo.UpdateMaas(id, model)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(model, nil)
}

func (m *ModelService) DeleteMaas(id int) *result.ErrorResult {
	err := m.repo.DeleteMaas(id)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(true, nil)
}

func ProvideMaasService(repo IModelRepo, messageService message.IMessageService, conversationService conversation.IConversationService) IMaasService {
	return &ModelService{repo: repo, messageService: messageService, conversationService: conversationService}
}

func (m *ModelService) Completion(c *gin.Context, req CompletionResponse, userId int64) {
	exist, err := m.conversationService.IsConversationExist(context.Background(), req.SessionId)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	if !exist {
		ret := utils.ResultWrapper(c)(nil, "conversation sessionId not exist")(utils.Error)
		c.JSON(400, ret)
		return
	}

	modelDetail, err := m.repo.FindMaasById(req.MaasId, &Maas{})
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	apiKey := modelDetail.ApiKey
	baseURL := modelDetail.BaseUrl
	model := modelDetail.Model
	prompt := req.Prompt
	params := req.Params
	sessionId := req.SessionId
	maasId := req.MaasId
	messageContext := req.Context

	// SequenceNumber 有bug
	maxSeq, err := m.messageService.GetMaxSequenceNumber(req.SessionId, message.MessageType(message.MessageTypes["Assistant"]))
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	// 创建问题
	question := message.NewMessage(
		message.WithConversationSessionID(sessionId),
		message.WithMessageType(message.MessageTypes["User"]),
		message.WithContent(prompt),
		message.WithMaasID(&maasId),
		message.WithSequenceNumber(maxSeq+1),
		message.WithContext(messageContext),
		message.WithParentQuestionID(nil),
	)

	currentQuestionMessage, err := m.messageService.CreateQuestionMessage(question)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	client := llm.NewClient()
	responseChan := make(chan string)

	go func() {
		err := client.SendQuery(apiKey, baseURL, model, prompt, params, responseChan)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
			c.JSON(400, ret)
			return
		}
	}()

	var fullResponse string

	// 使用 c.Stream 来确保连接保持开启状态直到所有数据发送完成
	c.Stream(func(w io.Writer) bool {
		chunk, ok := <-responseChan
		if !ok {
			return false
		}

		c.Writer.WriteString(chunk + "\n\n")
		c.Writer.Flush()

		// 如果是 DONE 消息，结束流
		if chunk == "data: [DONE]" {
			return false
		}

		content, err := client.GetContentFromChunk(chunk)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return false
		}
		fullResponse += content

		return true
	})

	parentQuestionID := currentQuestionMessage.ID
	// 保存回答数据
	answer := message.NewMessage(
		message.WithConversationSessionID(sessionId),
		message.WithMessageType(message.MessageTypes["Assistant"]),
		message.WithContent(fullResponse),
		message.WithMaasID(&maasId),
		message.WithSequenceNumber(maxSeq+2),
		message.WithContext(messageContext),
		message.WithParentQuestionID(&parentQuestionID),
	)
	// 更新conversation 最后一条信息的时间；
	err = m.conversationService.CreateModelAnswerMessage(answer)
	if err != nil {
		//ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		//c.JSON(400, ret)
		fmt.Println("err: ", err)
		return
	}

}

func (m *ModelService) InvokeMaas(ctx context.Context, req LLMRequest) (<-chan string, error) {
	client := llm.NewClient()
	responseChan := make(chan string)

	modelDetail, err := m.repo.FindMaasById(req.ModelId, &Maas{})
	if err != nil {
		return nil, err
	}

	fmt.Print("maas========: ", modelDetail)

	apiKey := modelDetail.ApiKey
	baseURL := modelDetail.BaseUrl
	model := modelDetail.Model
	prompt := req.Prompt
	params := req.Params

	go func() {
		err := client.SendQuery(apiKey, baseURL, model, prompt, params, responseChan)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	return responseChan, nil
}
