package model

type ModelAttrFunc func(u *Model)

type ModelAttrFuncs []ModelAttrFunc

func WithModel(model string) ModelAttrFunc {
	return func(u *Model) {
		u.Model = model
	}
}

func WithBaseUrl(baseUrl string) ModelAttrFunc {
	return func(u *Model) {
		u.BaseUrl = baseUrl
	}
}

func WithApiKey(apiKey string) ModelAttrFunc {
	return func(u *Model) {
		u.ApiKey = apiKey
	}
}

func WithModelType(modelType string) ModelAttrFunc {
	return func(u *Model) {
		u.ModelType = modelType
	}
}

func WithName(name string) ModelAttrFunc {
	return func(u *Model) {
		u.Name = name
	}
}

func WithAvatar(avatar string) ModelAttrFunc {
	return func(u *Model) {
		u.Avatar = avatar
	}
}

func WithCreateBy(createBy int64) ModelAttrFunc {
	return func(u *Model) {
		u.CreateBy = createBy
	}
}

func (this ModelAttrFuncs) apply(u *Model) {
	for _, f := range this {
		f(u)
	}
}
