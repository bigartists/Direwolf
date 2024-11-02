package user

import "strconv"

func ConvertUserToDTO(user *User) *UserBaseInfo {
	if user == nil {
		return nil
	}
	println("user=", user)
	return &UserBaseInfo{
		ID:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Avatar:      user.Avatar,
		Description: user.Description,
		CreateAt:    user.CreateTime,
		Admin:       strconv.Itoa(user.Admin),
	}
}
