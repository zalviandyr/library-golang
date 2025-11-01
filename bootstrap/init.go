package bootstrap

import (
	"library-be/app/controller"
	"library-be/config"
)

type Initialization struct {
	AuthorController    *controller.AuthorController
	PublisherController *controller.PublisherController
	BookController      *controller.BookController
	AuthController      *controller.AuthController
	Env                 config.Environment
}

func NewInitialization(
	AuthorController *controller.AuthorController,
	PublisherController *controller.PublisherController,
	BookController *controller.BookController,
	AuthController *controller.AuthController,
	Env config.Environment,
) *Initialization {
	return &Initialization{
		AuthorController:    AuthorController,
		PublisherController: PublisherController,
		BookController:      BookController,
		AuthController:      AuthController,
		Env:                 Env,
	}
}
