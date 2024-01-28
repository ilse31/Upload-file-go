package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	helloController := controllers.NewHelloController()

	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Get("/hello", helloController.Show)
}
