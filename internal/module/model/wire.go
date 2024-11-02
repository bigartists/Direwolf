package model

import "github.com/google/wire"

var Modelset = wire.NewSet(
	ProvideModelRepo,
	ProvideModelService,
	ProvideModelController,
)
