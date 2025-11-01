package bootstrap

import "library-be/app/controller"

type Initialization struct {
	AuthorController    *controller.AuthorController
	PublisherController *controller.PublisherController
	BookController      *controller.BookController
	AuthController      *controller.AuthController
}

func NewInitialization(
	AuthorController *controller.AuthorController,
	PublisherController *controller.PublisherController,
	BookController *controller.BookController,
	AuthController *controller.AuthController,
) *Initialization {
	return &Initialization{
		AuthorController:    AuthorController,
		PublisherController: PublisherController,
		BookController:      BookController,
		AuthController:      AuthController,
	}
}
