package user

import (
	"fmt"
	"github.com/bigartists/Direwolf/pkg/result"
)

type IUserService interface {
	GetUserList() []*UserBaseInfo
	GetUserDetail(id int64) *result.ErrorResult

	UpdateUser(id int, user *User) *result.ErrorResult
	DeleteUser(id int) *result.ErrorResult
	SignIn(username string, password string) (*UserBaseInfo, error)
	Register(email string, username string, password string) error
}

type UserService struct {
	repo IUserRepo
}

func ProvideUserService(repo IUserRepo) IUserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignIn(username string, password string) (*UserBaseInfo, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	//if user.Password != password {
	//	err = fmt.Errorf("用户名%s或密码错误", username)
	//	return nil, err
	//}
	// 校验密码
	if !user.CheckPassword(password) {
		err = fmt.Errorf("用户名%s或密码错误", username)
		return nil, err
	}
	userdto := ConvertUserToDTO(user)
	return userdto, nil
}

func (s *UserService) Register(email string, username string, password string) error {

	if _, err := s.repo.FindUserByUsername(username); err == nil {
		fmt.Println("errerr===", err)
		fmt.Println("用户名%s已存在", username, err)
		return fmt.Errorf("用户名%s已存在", username)
	}

	if _, err := s.repo.FindUserByEmail(email); err == nil {
		return fmt.Errorf("邮箱%s已存在", email)
	}

	user := NewModel(WithEmail(email), WithUsername(username), WithPassword(password), WithActive(new(bool)))

	*user.Active = true

	err := user.GeneratePassword()
	if err != nil {
		return fmt.Errorf("密码加密失败")
	}
	err = s.repo.CreateUser(user)
	fmt.Printf("user:%v\n", user)
	if err != nil {
		return fmt.Errorf("用户注册失败")
	}

	return nil
}

func (s *UserService) GetUserList() []*UserBaseInfo {
	users := s.repo.FindUserAll()
	userdtos := make([]*UserBaseInfo, len(users))
	for i, user := range users {
		userdtos[i] = ConvertUserToDTO(user)
	}
	return userdtos
}

func (s *UserService) GetUserDetail(id int64) *result.ErrorResult {
	//TODO implement me
	user := NewModel()
	_, err := s.repo.FindUserById(id, user)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(user, nil)
}

func (s *UserService) UpdateUser(id int, user *User) *result.ErrorResult {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) DeleteUser(id int) *result.ErrorResult {
	//TODO implement me
	panic("implement me")
}

//
//// 更新用户
//func (this *GetterImpl) UpdateUser(id int, user *User.UserModelImpl) *result.ErrorResult {
//	db := dbs.Orm.Where("id=?", id).Updates(user)
//	if db.Error != nil {
//		return result.Result(nil, db.Error)
//	}
//	return result.Result(user, nil)
//}
//
//// 删除用户
//func (this *GetterImpl) DeleteUser(id int) *result.ErrorResult {
//	user := User.New()
//	db := dbs.Orm.Where("id=?", id).Delete(user)
//	if db.Error != nil {
//		return result.Result(nil, db.Error)
//	}
//	return result.Result(user, nil)
//}
