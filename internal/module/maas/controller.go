package maas

import (
	"github.com/bigartists/Direwolf/internal/ErrorMessage"
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/bigartists/Direwolf/pkg/result"
	"github.com/bigartists/Direwolf/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

type MaasController struct {
	service IMaasService
}

func ProvideMaasController(service IMaasService) *MaasController {
	return &MaasController{service: service}
}

func (this *MaasController) Create(c *gin.Context) {
	var model MaasRequest
	result.Result(c.ShouldBindJSON(&model)).Unwrap()

	userId := user.GetUserId(c)

	if userId == 0 {
		ret := utils.ResultWrapper(c)(nil, ErrorMessage.InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}

	createModel, err := this.service.ImportMaas(&model, userId)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	c.JSON(200, createModel)
}

func (this *MaasController) GetMaasList(c *gin.Context) {
	userId := user.GetUserId(c)

	if userId == 0 {
		ret := utils.ResultWrapper(c)(nil, ErrorMessage.InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}
	list, err := this.service.GetMaasList(userId)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	ret := utils.ResultWrapper(c)(list, "")(utils.OK)
	c.JSON(200, ret)
}

func (this *MaasController) Update(c *gin.Context) {
	var model Maas
	result.Result(c.ShouldBindJSON(&model)).Unwrap()
	result.Result(this.service.UpdateMaas(int(model.ID), &model)).Unwrap()
	c.JSON(200, model)
}

func (this *MaasController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	result.Result(this.service.DeleteMaas(id)).Unwrap()
	c.JSON(200, gin.H{"id": id})
}

func (this *MaasController) GetMaasDetail(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	result.Result(this.service.GetMaasDetail(id)).Unwrap()
	c.JSON(200, gin.H{"id": id})
}

type LLMRequest struct {
	ModelId int64                  `json:"model_id" binding:"required"`
	Prompt  string                 `json:"prompt" binding:"required"`
	Params  map[string]interface{} `json:"params" binding:"required"`
}

type CompletionResponse struct {
	ModelId   int64                  `json:"model_id" binding:"required"`
	Prompt    string                 `json:"prompt" binding:"required"`
	Params    map[string]interface{} `json:"params" binding:"required"`
	SessionId string                 `json:"session_id" binding:"required"`
	Context   string                 `json:"context" binding:"required"`
}

func (this *MaasController) Invoke(c *gin.Context) {
	var req LLMRequest
	result.Result(c.ShouldBindJSON(&req)).Unwrap()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	responseChan, err := this.service.InvokeMaas(c.Request.Context(), req)
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

func (this *MaasController) Completion(c *gin.Context) {
	var req CompletionResponse
	result.Result(c.ShouldBindJSON(&req)).Unwrap()

	userId := user.GetUserId(c)

	if userId == 0 {
		ret := utils.ResultWrapper(c)(nil, ErrorMessage.InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}
	this.service.Completion(c, req, userId)
}

func (this *MaasController) Build(r *gin.RouterGroup) {
	r.POST("/maas/create", this.Create)
	r.POST("/maas/update", this.Update)
	r.POST("/maas/delete/:id", this.Delete)

	r.GET("/maas/:id", this.GetMaasDetail)
	r.GET("/maas_list", this.GetMaasList)
	r.POST("/maas/invoke", this.Invoke)
	r.POST("/maas/completions", this.Completion)
}
