package controllers

import (
	"net/http"

	"github.com/victorananias/go-auth-api/repositories"
	"github.com/victorananias/go-auth-api/requests"
	"github.com/victorananias/go-auth-api/responses"
	"github.com/victorananias/go-auth-api/services"
)

type authController struct {
	*baseController
	authService    *services.AuthService
	userRepository *repositories.UserRepository
}

var AuthController *authController = NewAuthController()

func NewAuthController() *authController {
	controller := &authController{}
	controller.authService = services.NewAuthService()
	controller.userRepository = repositories.NewUserRepository()
	return controller
}

func (controller *authController) Register(response http.ResponseWriter, request *http.Request) {
	var user requests.RegisterRequest
	err := controller.decodeRequest(request, &user)
	if err != nil {
		controller.respondWithError(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = controller.authService.Register(user)
	if err != nil {
		controller.respondWithError(response, err.Error(), http.StatusBadRequest)
		return
	}
	controller.respondWithSuccess(response, "Registered.", http.StatusCreated)
}

func (controller *authController) Login(response http.ResponseWriter, request *http.Request) {
	var loginRequest requests.LoginRequest
	err := controller.decodeRequest(request, &loginRequest)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := controller.authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		controller.respondWithError(response, err.Error(), http.StatusUnauthorized)
		return
	}
	loginResponse := responses.LoginResponse{Token: token}
	controller.respondWithJson(response, loginResponse, http.StatusOK)
}

func (controller *authController) List(response http.ResponseWriter, request *http.Request) {
	users, err := controller.userRepository.List()
	if err != nil {
		controller.respondWithError(response, err.Error(), http.StatusBadRequest)
		return
	}
	controller.respondWithJson(response, users, http.StatusOK)
}

// func (controller *authController) ValidateToken(response http.ResponseWriter, request *http.Request) {
// 	jwt := services.Jwt{}
// 	var validateTokenRequest requests.ValidateTokenRequest
// 	err := json.NewDecoder(request.Body).Decode(&validateTokenRequest)
// 	if err != nil {
// 		http.Error(response, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	isValid := jwt.Validade(validateTokenRequest.Token)
// 	controller.respondWithJson(response, responses.ValidateTokenResponse{Valid: isValid}, http.StatusOK)
// }
