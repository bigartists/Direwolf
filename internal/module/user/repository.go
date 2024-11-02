package user

import (
	"fmt"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindUserAll() []*User
	FindUserById(id int64, user *User) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(id int, user *User) error
	DeleteUser(id int) error
}

type UserRepo struct {
	db *gorm.DB
}

func ProvideUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindUserByUsername(username string) (*User, error) {
	var user User
	err := u.db.Where("username=?", username).First(&user).Error
	return &user, err
}

func (u *UserRepo) FindUserByEmail(email string) (*User, error) {
	var user User
	err := u.db.Where("email=?", email).First(&user).Error
	return &user, err
}

func (u *UserRepo) CreateUser(user *User) error {
	return u.db.Create(user).Error
}

func (u *UserRepo) UpdateUser(id int, user *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) FindUserById(id int64, user *User) (*User, error) {
	//TODO implement me
	db := u.db.Where("id=?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found, id=%d", id)
	}
	return user, nil
}

func (u *UserRepo) FindUserAll() []*User {
	var users []*User
	u.db.Find(&users)
	return users
}
