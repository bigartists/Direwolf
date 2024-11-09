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
	"time"
)

type MaasBaseInfo struct {
	Id       int64     `json:"id"`
	Model    string    `json:"maas"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_time"`
}

type IMaasService interface {
	GetMaasList(userId int64) ([]*Maas, error)
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

func (m *ModelService) GetMaasList(userId int64) ([]*Maas, error) {
	modelList, err := m.repo.FindMaasList(userId)
	//var modelBaseInfoList []*MaasBaseInfo
	//for _, maas := range modelList {
	//	modelBaseInfoList = append(modelBaseInfoList, &MaasBaseInfo{
	//		Id:       maas.ID,
	//		Model:    maas.Model,
	//		Avatar:   maas.Avatar,
	//		CreateAt: maas.CreateTime,
	//	})
	//}
	//return modelBaseInfoList
	return modelList, err
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

	modelDetail, err := m.repo.FindMaasById(req.ModelId, &Maas{})
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
	modelId := req.ModelId
	messageContext := req.Context

	maxSeq, err := m.messageService.GetMaxSequenceNumber(req.SessionId)
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
		message.WithModelID(&modelId),
		message.WithSequenceNumber(maxSeq+1),
		message.WithContext(messageContext),
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

	for chunk := range responseChan {
		c.SSEvent("", chunk)
		c.Writer.Flush()
		content, err := client.GetContentFromChunk(chunk)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
			c.JSON(400, ret)
			return
		}
		fullResponse += content
	}

	// 保存回答数据
	answer := message.NewMessage(
		message.WithConversationSessionID(sessionId),
		message.WithMessageType(message.MessageTypes["Assistant"]),
		message.WithContent(fullResponse),
		message.WithModelID(&modelId),
		message.WithSequenceNumber(maxSeq+2),
		message.WithContext(messageContext),
		message.WithParentQuestionID(currentQuestionMessage.ID),
	)
	// 更新conversation 最后一条信息的时间；
	err = m.messageService.CreateModelAnswerMessage(answer)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}

	println("fullResponse: ", fullResponse)
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
