package main

import (
	"library-be/app/router"
	"library-be/bootstrap"
)

func main() {
	init := bootstrap.Initialized()
	app := router.Init(init)

	app.Run()
}
