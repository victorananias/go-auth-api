package main

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string
	Password string
}

type DefaultResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

var database Database

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		respondWithError(responseWriter, "Error while registering.", http.StatusBadRequest)
	}

	database.CreateUser(user)

	respondWithSuccess(responseWriter, "Registered.", http.StatusCreated)
}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var loginRequest LoginRequest

		err := json.NewDecoder(request.Body).Decode(&loginRequest)

		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}

		loginRequest.Password = hashPassword(loginRequest.Password)
		respondWithJson(responseWriter, loginRequest, http.StatusOK)
	}
}

func respondWithError(responseWriter http.ResponseWriter, message string, status int) {
	response := DefaultResponse{Message: message, Success: false}
	respondWithJson(responseWriter, response, status)
}

func respondWithSuccess(responseWriter http.ResponseWriter, message string, status int) {
	response := DefaultResponse{Message: message, Success: true}
	respondWithJson(responseWriter, response, status)
}

func respondWithJson(responseWriter http.ResponseWriter, i interface{}, status int) {
	responseWriter.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(i)
	responseWriter.WriteHeader(status)
	responseWriter.Write(jsonResponse)
}
