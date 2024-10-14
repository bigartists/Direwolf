package user

import (
	"fmt"
	"github.com/bigartists/Direwolf/config"
	"github.com/bigartists/Direwolf/pkg/result"
	. "github.com/bigartists/Direwolf/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (this *UserController) Login(c *gin.Context) {
	// 校验输入参数是否合法
	params := &struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	user, err := ServiceGetter.SignIn(params.Username, params.Password)
	if err != nil {
		ret := ResultWrapper(c)(nil, err.Error())(Error)
		c.JSON(400, ret)
		return
	}

	//// 生成 token
	prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
	curTime := time.Now().Add(time.Second * 60 * 60 * 24)
	token, _ := GenerateToken(user.ID, prikey, curTime)

	c.Set("token", token)
	ret := ResultWrapper(c)(user, "")(OK)
	c.JSON(200, ret)
}

//func (this *UserController) SignUp(c *gin.Context) {
//	// 校验输入参数是否合法
//	params := &dto.SignupRequest{}
//	// 校验参数
//	result.Result(c.ShouldBindJSON(params)).Unwrap()
//
//	err := ServiceGetter.SignUp(params.Email, params.Username, params.Password)
//	if err != nil {
//		ret := ResultWrapper(c)(nil, err.Error())(Error)
//		c.JSON(400, ret)
//	}
//	ret := ResultWrapper(c)(true, "")(Created)
//	c.JSON(201, ret)
//}

func GetAuthUser(c *gin.Context) *User {
	t, exist := c.Get("auth_user")
	if !exist {
		return nil
	}
	return t.(*User)
}

func (this *UserController) UserList(c *gin.Context) {
	ret := ResultWrapper(c)(ServiceGetter.GetUserList(), "")(OK)
	c.JSON(200, ret)
}

func (this *UserController) UserDetail(c *gin.Context) {
	id := &struct {
		Id int64 `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	ret := ResultWrapper(c)(ServiceGetter.GetUserDetail(id.Id).Unwrap(), "")(OK)
	c.JSON(200, ret)
}

func (this *UserController) GetMe(c *gin.Context) {
	u := GetAuthUser(c)
	ret := ResultWrapper(c)(u, "")(OK)
	c.JSON(200, ret)
}

type LLMRequest struct {
	BaseURL string                 `json:"base_url" binding:"required"`
	APIKey  string                 `json:"api_key" binding:"required"`
	Model   string                 `json:"model" binding:"required"`
	Prompt  string                 `json:"prompt" binding:"required"`
	Params  map[string]interface{} `json:"params" binding:"required"`
}

func (this *UserController) Invoke(c *gin.Context) {
	var req LLMRequest
	result.Result(c.ShouldBindJSON(&req)).Unwrap()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	responseChan, err := ServiceGetter.InvokeModel(c.Request.Context(), req)
	if err != nil {
		c.SSEvent("error", err.Error())
		return
	}

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-responseChan; ok {
			fmt.Println("msg: ", msg)
			c.SSEvent("", msg)
			return true
		}
		return false
	})
}

func (this *UserController) Build(r *gin.RouterGroup) {
	r.POST("/login", this.Login)
	//r.POST("/register", this.SignUp)
	r.GET("/users", this.UserList)
	r.GET("/user/:id", this.UserDetail)
	r.GET("/me", this.GetMe)
	r.POST("/invoke", this.Invoke)
}
