package model

import (
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/bigartists/Direwolf/pkg/result"
	"github.com/bigartists/Direwolf/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

type ModelController struct {
	service IModelService
}

func ProvideModelController(service IModelService) *ModelController {
	return &ModelController{service: service}
}

func (this *ModelController) Create(c *gin.Context) {
	var model ModelRequest
	result.Result(c.ShouldBindJSON(&model)).Unwrap()

	userId := user.GetUserId(c)

	if userId == 0 {
		ret := utils.ResultWrapper(c)(nil, InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}

	createModel, err := this.service.CreateModel(&model, userId)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	c.JSON(200, createModel)
}

func (this *ModelController) GetModelList(c *gin.Context) {
	userId := user.GetUserId(c)

	if userId == 0 {
		ret := utils.ResultWrapper(c)(nil, InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}
	list, err := this.service.GetModelList(userId)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	ret := utils.ResultWrapper(c)(list, "")(utils.OK)
	c.JSON(200, ret)
}

func (this *ModelController) Update(c *gin.Context) {
	var model Model
	result.Result(c.ShouldBindJSON(&model)).Unwrap()
	result.Result(this.service.UpdateModel(int(model.ID), &model)).Unwrap()
	c.JSON(200, model)
}

func (this *ModelController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	result.Result(this.service.DeleteModel(id)).Unwrap()
	c.JSON(200, gin.H{"id": id})
}

func (this *ModelController) GetModelDetail(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	result.Result(this.service.GetModelDetail(id)).Unwrap()
	c.JSON(200, gin.H{"id": id})
}

type LLMRequest struct {
	BaseURL string                 `json:"base_url" binding:"required"`
	APIKey  string                 `json:"api_key" binding:"required"`
	Model   string                 `json:"model" binding:"required"`
	Prompt  string                 `json:"prompt" binding:"required"`
	Params  map[string]interface{} `json:"params" binding:"required"`
}

func (this *ModelController) Invoke(c *gin.Context) {
	var req LLMRequest
	result.Result(c.ShouldBindJSON(&req)).Unwrap()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	responseChan, err := this.service.InvokeModel(c.Request.Context(), req)
	if err != nil {
		c.SSEvent("error", err.Error())
		return
	}

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-responseChan; ok {
			c.SSEvent("", msg)
			return true
		}
		return false
	})
}

func (this *ModelController) Build(r *gin.RouterGroup) {
	r.POST("/model/create", this.Create)
	r.POST("/model/update", this.Update)
	r.POST("/model/delete/:id", this.Delete)

	r.GET("/model/:id", this.GetModelDetail)
	r.GET("/models", this.GetModelList)
	r.POST("/model/invoke", this.Invoke)
}
