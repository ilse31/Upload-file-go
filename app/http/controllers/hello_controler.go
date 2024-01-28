package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type HelloController struct {
	//Dependent services
}

func NewHelloController() *HelloController {
	return &HelloController{
		//Inject services
	}
}

func (r *HelloController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})

}
