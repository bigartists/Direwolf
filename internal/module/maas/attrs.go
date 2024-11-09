package maas

type MaasAttrFunc func(u *Maas)

type MaasAttrFuncs []MaasAttrFunc

func WithModel(model string) MaasAttrFunc {
	return func(u *Maas) {
		u.Model = model
	}
}

func WithBaseUrl(baseUrl string) MaasAttrFunc {
	return func(u *Maas) {
		u.BaseUrl = baseUrl
	}
}

func WithApiKey(apiKey string) MaasAttrFunc {
	return func(u *Maas) {
		u.ApiKey = apiKey
	}
}

func WithModelType(modelType string) MaasAttrFunc {
	return func(u *Maas) {
		u.ModelType = modelType
	}
}

func WithName(name string) MaasAttrFunc {
	return func(u *Maas) {
		u.Name = name
	}
}

func WithAvatar(avatar string) MaasAttrFunc {
	return func(u *Maas) {
		u.Avatar = avatar
	}
}

func WithCreateBy(createBy int64) MaasAttrFunc {
	return func(u *Maas) {
		u.CreateBy = createBy
	}
}

func (this MaasAttrFuncs) apply(u *Maas) {
	for _, f := range this {
		f(u)
	}
}
