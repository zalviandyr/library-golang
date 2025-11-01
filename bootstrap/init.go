package bootstrap

import "library-be/app/controller"

type Initialization struct {
	AuthorController    *controller.AuthorController
	PublisherController *controller.PublisherController
}

func NewInitialization(
	AuthorController *controller.AuthorController,
	PublisherController *controller.PublisherController,
) *Initialization {
	return &Initialization{
		AuthorController:    AuthorController,
		PublisherController: PublisherController,
	}
}
