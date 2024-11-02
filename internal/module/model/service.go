package model

import (
	"context"
	"fmt"
	"github.com/bigartists/Direwolf/pkg/llm"
	"github.com/bigartists/Direwolf/pkg/result"
	"time"
)

type ModelBaseInfo struct {
	Id       int64     `json:"id"`
	Model    string    `json:"model"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_time"`
}

type IModelService interface {
	GetModelList(userId int64) ([]*Model, error)
	GetModelDetail(id int64) *result.ErrorResult
	CreateModel(model *ModelRequest, userId int64) (*Model, error)
	UpdateModel(id int, model *Model) *result.ErrorResult
	DeleteModel(id int) *result.ErrorResult
	InvokeModel(ctx context.Context, req LLMRequest) (<-chan string, error)
}

type ModelService struct {
	repo IModelRepo
}

func (m *ModelService) GetModelList(userId int64) ([]*Model, error) {
	modelList, err := m.repo.FindModelList(userId)
	//var modelBaseInfoList []*ModelBaseInfo
	//for _, model := range modelList {
	//	modelBaseInfoList = append(modelBaseInfoList, &ModelBaseInfo{
	//		Id:       model.ID,
	//		Model:    model.Model,
	//		Avatar:   model.Avatar,
	//		CreateAt: model.CreateTime,
	//	})
	//}
	//return modelBaseInfoList
	return modelList, err
}

func (m *ModelService) GetModelDetail(id int64) *result.ErrorResult {
	model := &Model{}
	model, err := m.repo.FindModelById(id, model)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(model, nil)
}

func (m *ModelService) CreateModel(modelRequest *ModelRequest, userId int64) (*Model, error) {
	isDuplicate, err := m.repo.CheckDuplicateModelBaseURL(context.Background(), modelRequest.Model, modelRequest.BaseURL)

	fmt.Println("isDuplicate:= ", isDuplicate)
	fmt.Println("err:= ", err)

	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, ErrDuplicateModelBaseURL
	}
	fmt.Println(".................=======")
	model := NewModel(
		WithModel(modelRequest.Model),
		WithApiKey(modelRequest.APIKey),
		WithBaseUrl(modelRequest.BaseURL),
		WithModelType(modelRequest.ModelType),
		WithName(modelRequest.Name),
		WithCreateBy(userId),
	)

	err = m.repo.CreateModel(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *ModelService) UpdateModel(id int, model *Model) *result.ErrorResult {
	err := m.repo.UpdateModel(id, model)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(model, nil)
}

func (m *ModelService) DeleteModel(id int) *result.ErrorResult {
	err := m.repo.DeleteModel(id)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(true, nil)
}

func ProvideModelService(repo IModelRepo) IModelService {
	return &ModelService{repo: repo}
}

func (m *ModelService) InvokeModel(ctx context.Context, req LLMRequest) (<-chan string, error) {
	client := llm.NewClient()
	responseChan := make(chan string)

	go func() {
		err := client.SendQuery(req.APIKey, req.BaseURL, req.Model, req.Prompt, req.Params, responseChan)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	return responseChan, nil
}
