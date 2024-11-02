package user

import "github.com/google/wire"

var Userset = wire.NewSet(
	ProvideUserRepo,
	ProvideUserService,
	ProvideUserController,
)
