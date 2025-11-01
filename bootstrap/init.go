package bootstrap

import "library-be/app/controller"

type Initialization struct {
	AuthorController *controller.AuthorController
}

func NewInitialization(
	AuthorController *controller.AuthorController,
) *Initialization {
	return &Initialization{
		AuthorController: AuthorController,
	}
}
