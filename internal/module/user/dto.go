package user

import "time"

// UserBaseInfo 表示输出到外部的用户信息

type UserBaseInfo struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Admin       string    `json:"admin"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"created_at"`
}

type UserAsCreator struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
