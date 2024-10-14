package vars

import "net/http"

const (
	ContextKeyToken         = "token"
	HTTPSUCCESS             = http.StatusOK
	HTTPFAIL                = http.StatusBadRequest
	HTTPUNAUTHORIZED        = http.StatusUnauthorized
	HTTPMESSAGESUCCESS      = "success"
	HTTPMESSAGEFAIL         = "fail"
	HTTPMESSAGEUNAUTHORIZED = "Unauthorized"

	STAR_TARGET_PROJECT = "project"
)
