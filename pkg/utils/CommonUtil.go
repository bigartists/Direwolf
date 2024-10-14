package utils

import (
	"github.com/bigartists/Direwolf/pkg/vars"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
)

func FormatTime(str v1.Time) string {
	return str.Format("2006-01-02 15:04:05")
}

type ResultHeader struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

type JSONResult struct {
	Payload interface{}   `json:"payload"`
	Header  *ResultHeader `json:"header"`
}

func NewJSONResult(result interface{}) *JSONResult {
	header := &ResultHeader{
		Code:  0,
		Msg:   "",
		Token: "",
	}
	return &JSONResult{
		Header:  header,
		Payload: result,
	}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult(nil)
		},
	}
}

type Output func(c *gin.Context, v interface{}) interface{}

type ResultFunc func(result interface{}, message string) func(output Output) interface{}

func ResultWrapper(c *gin.Context) ResultFunc {
	return func(result interface{}, message string) func(output Output) interface{} {
		r := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(r)

		token := c.GetString("token")

		r.Header.Msg = message
		r.Header.Token = token
		r.Payload = result

		//r.Result = map[string]interface{}{
		//	"data":  result,
		//	"token": token,
		//}
		return func(output Output) interface{} {
			return output(c, r)
		}
	}
}

func OK(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {

		r.Header.Code = vars.HTTPSUCCESS
		r.Header.Msg = vars.HTTPMESSAGESUCCESS
		return r
	}
	return nil
}

func Created(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Header.Code = vars.HTTPSUCCESS
		r.Header.Msg = vars.HTTPMESSAGESUCCESS
		return r
	}
	return nil
}

func Error(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Header.Code = vars.HTTPFAIL
		if r.Header.Msg == "" {
			r.Header.Msg = vars.HTTPMESSAGEFAIL
		}
		return r
	}
	return nil
}

func Unauthorized(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Header.Code = vars.HTTPUNAUTHORIZED
		if r.Header.Msg == "" {
			r.Header.Msg = vars.HTTPMESSAGEUNAUTHORIZED
		}
		return r
	}
	return nil
}
