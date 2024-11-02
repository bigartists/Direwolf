//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/bigartists/Direwolf/client"
	"github.com/bigartists/Direwolf/internal"
	"github.com/bigartists/Direwolf/internal/routers"
	"github.com/bigartists/Direwolf/pkg/middlewares"
	"github.com/bigartists/Direwolf/server"
	"github.com/google/wire"
)

func InitializeApp() (*server.App, error) {
	wire.Build(
		client.ProvideDB,
		internal.WholeSets,
		middlewares.NewAuthMiddleware,
		routers.ProvideRouter,
		server.ProvideApp,
	)
	return &server.App{}, nil
}
