package main

import (
	"net/http"

	"github.com/victorananias/go-auth-api/controllers"
)

type Routes struct {
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (routes *Routes) Register() {
	routes.register(http.MethodPost, "/", controllers.AuthController.List)
	routes.register(http.MethodPost, "/register", controllers.AuthController.Register)
	routes.register(http.MethodPost, "/login", controllers.AuthController.Login)
	// routes.register(http.MethodPost, "/validate-token", controllers.AuthController.ValidateToken)
}

func (routes *Routes) register(method string, route string, handler HandlerFunc) {
	http.HandleFunc(route, func(response http.ResponseWriter, request *http.Request) {
		routes.enableCors(response)
		if request.Method == method {
			handler(response, request)
		}
	})
}

func (routes *Routes) enableCors(response http.ResponseWriter) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
