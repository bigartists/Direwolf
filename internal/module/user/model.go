package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 创建 Users struct
type User struct {
	Id int64 `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	//Email 邮箱，不能为空， 必须是邮箱格式，且不能重复；
	Email string `json:"email" gorm:"column:email;unique" binding:"required,email"`
	//Username 用户名，创建自定义验证器： 长度在 6-20 之间，且不能重复，只能包含大小写字母，数字，下划线；第一个字符必须是字母；
	Username string `json:"username" gorm:"column:username;unique" binding:"usernameValid"`
	//Password 密码， 长度在 6-20 之间，只能包含字母，数字，下划线；
	Password string `json:"-" gorm:"column:password" binding:"passwordValid"`
	// admin 0 和 1
	Admin int `json:"admin" gorm:"column:admin"`
	// active 0 和 1

	Active *bool `json:"active" gorm:"column:active"`
	// nickname
	Name string `json:"name" gorm:"column:name"`
	// description
	Description string `json:"description" gorm:"column:description"`
	// avatar
	Avatar string `json:"avatar" gorm:"column:avatar"`
	// 自动维护时间

	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
	CreateBy   int64     `json:"-" gorm:"column:create_by"`
	UpdateBy   int64     `json:"-" gorm:"column:update_by"`
}

func (u *User) TableName() string {
	return "user"
}

func NewModel(attrs ...UserModelAttrFunc) *User {
	u := &User{}
	UserModelAttrFuncs(attrs).apply(u)
	return u
}

func (u *User) Mutate(attrs ...UserModelAttrFunc) *User {
	UserModelAttrFuncs(attrs).apply(u)
	return u
}

// 生成密码
func (u *User) GeneratePassword() error {
	// 使用 bcrypt 生成密码, bcrypt.DefaultCost 表示默认的加密强度，值越大加密强度越大，但是会消耗更多的资源
	pas, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pas)
	return nil
}

// 检查密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

//func (u *User) BeforeSave() error {
//	//turn password into hash
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//	u.Password = string(hashedPassword)
//	return nil
//}
