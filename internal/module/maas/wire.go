package maas

import "github.com/google/wire"

var MaasSet = wire.NewSet(
	ProvideMaasRepo,
	ProvideMaasService,
	ProvideMaasController,
)
