package user

import (
	"context"
	"github.com/bigartists/Direwolf/pkg/result"
)

type Service interface {
	GetUserList() []*UserBaseInfo
	GetUserDetail(id int64) *result.ErrorResult
	CreateUser(user *User) *result.ErrorResult
	UpdateUser(id int, user *User) *result.ErrorResult
	DeleteUser(id int) *result.ErrorResult
	SignIn(username string, password string) (*UserBaseInfo, error)
	SignUp(email string, username string, password string) error
	InvokeModel(ctx context.Context, req LLMRequest) (<-chan string, error)
}
