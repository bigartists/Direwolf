package user

type UserModelAttrFunc func(u *User)

type UserModelAttrFuncs []UserModelAttrFunc

func WithUserId(id int64) UserModelAttrFunc {
	return func(u *User) {
		u.Id = id
	}
}

func WithUsername(name string) UserModelAttrFunc {
	return func(u *User) {
		u.Username = name
	}
}

func WithPassword(pwd string) UserModelAttrFunc {
	return func(u *User) {
		u.Password = pwd
	}
}

func WithEmail(email string) UserModelAttrFunc {
	return func(u *User) {
		u.Email = email
	}
}

func WithActive(active *bool) UserModelAttrFunc {
	return func(u *User) {
		u.Active = active
	}
}

func (this UserModelAttrFuncs) apply(u *User) {
	for _, f := range this {
		f(u)
	}
}
