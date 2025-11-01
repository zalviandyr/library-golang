//go:build wireinject
// +build wireinject

package bootstrap

import (
	"library-be/app/controller"
	"library-be/app/repository"
	"library-be/app/service"
	"library-be/config"

	"github.com/google/wire"
)

var authorSet = wire.NewSet(controller.NewAuthorController, service.NewAuthorService, repository.NewAuthorRepository)

func Initialized() *Initialization {
	wire.Build(
		NewInitialization,
		config.InitDatabase,
		config.InitEnvironment,

		authorSet,
	)

	return nil
}
