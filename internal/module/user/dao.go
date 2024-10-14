package user

import (
	"fmt"
	"github.com/bigartists/Direwolf/client"
)

type IUserDao interface {
	FindUserAll() []*User
	FindUserById(id int64, user *User) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(id int, user *User) error
	DeleteUser(id int) error
}

var DaoGetter IUserDao

func init() {
	DaoGetter = NewGetterImpl()
}

type GetterImpl struct {
}

func NewGetterImpl() *GetterImpl {
	return &GetterImpl{}
}

func (I GetterImpl) FindUserByUsername(username string) (*User, error) {
	var user User
	err := client.Orm.Where("username=?", username).Find(&user).Error

	return &user, err
}

func (I GetterImpl) FindUserByEmail(email string) (*User, error) {
	var user User
	err := client.Orm.Where("email=?", email).Find(&user).Error
	return &user, err
}

func (I GetterImpl) CreateUser(user *User) error {
	return client.Orm.Create(user).Error
}

func (I GetterImpl) UpdateUser(id int, user *User) error {
	//TODO implement me
	panic("implement me")
}

func (I GetterImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (I GetterImpl) FindUserById(id int64, user *User) (*User, error) {
	//TODO implement me
	db := client.Orm.Where("id=?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found, id=%d", id)
	}
	return user, nil
}

func (I GetterImpl) FindUserAll() []*User {
	var users []*User
	client.Orm.Find(&users)
	return users
}
