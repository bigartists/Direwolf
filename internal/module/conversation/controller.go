package conversation

import (
	"github.com/bigartists/Direwolf/internal/ErrorMessage"
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/bigartists/Direwolf/pkg/result"
	"github.com/bigartists/Direwolf/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	service IConversationService
}

func ProvideConversationController(service IConversationService) *ConversationController {
	return &ConversationController{service: service}
}

func (this *ConversationController) Build(r *gin.RouterGroup) {
	r.POST("/conversation/create", this.Create)
	r.GET("/conversations", this.GetConversationHistory)
	r.GET("/conversation/:id", this.GetConversationDetail)
}

func (this *ConversationController) Create(c *gin.Context) {
	var req CreateConversationRequest

	result.Result(c.ShouldBindJSON(&req)).Unwrap()

	conversation, err := this.service.CreateConversation(c, req)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	ret := utils.ResultWrapper(c)(conversation, "")(utils.OK)
	c.JSON(200, ret)
}

func (this *ConversationController) GetConversationHistory(c *gin.Context) {
	userID := user.GetUserId(c)

	if userID == 0 {
		ret := utils.ResultWrapper(c)(nil, ErrorMessage.InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}
	history, err := this.service.GetConversationHistory(c, userID)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	ret := utils.ResultWrapper(c)(history, "")(utils.OK)
	c.JSON(200, ret)
}

func (this *ConversationController) GetConversationDetail(c *gin.Context) {
	userID := user.GetUserId(c)
	if userID == 0 {
		ret := utils.ResultWrapper(c)(nil, ErrorMessage.InValidUserId)(utils.Error)
		c.JSON(400, ret)
		return
	}
	conversationID := c.GetInt64("id")
	detail, err := this.service.GetConversationDetail(c, userID, conversationID)
	if err != nil {
		ret := utils.ResultWrapper(c)(nil, err.Error())(utils.Error)
		c.JSON(400, ret)
		return
	}
	ret := utils.ResultWrapper(c)(detail, "")(utils.OK)
	c.JSON(200, ret)
}