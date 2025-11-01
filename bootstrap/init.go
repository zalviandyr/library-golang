package bootstrap

import "library-be/app/controller"

type Initialization struct {
	AuthorController    *controller.AuthorController
	PublisherController *controller.PublisherController
	BookController      *controller.BookController
}

func NewInitialization(
	AuthorController *controller.AuthorController,
	PublisherController *controller.PublisherController,
	BookController *controller.BookController,
) *Initialization {
	return &Initialization{
		AuthorController:    AuthorController,
		PublisherController: PublisherController,
		BookController:      BookController,
	}
}
