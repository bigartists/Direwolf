package internal

import (
	"github.com/bigartists/Direwolf/internal/module/model"
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/google/wire"
)

var WholeSets = wire.NewSet(
	user.Userset,
	model.Modelset,
)
