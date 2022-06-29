package main

import (
	"encoding/json"
	"log"
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

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		return
	}
	log.Println("Register request called.")
	userRepository := newUserRepository()
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		respondWithError(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	err, _ = userRepository.Register(user)
	if err != nil {
		respondWithError(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithSuccess(responseWriter, "Registered.", http.StatusCreated)
}

func List(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		return
	}
	log.Println("List request called.")
	userRepository := newUserRepository()
	err, users := userRepository.List()
	if err != nil {
		respondWithError(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJson(responseWriter, users, http.StatusOK)
}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		return
	}
	log.Println("Login request called.")
	userRepository := newUserRepository()
	var loginRequest LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	if ok := userRepository.Login(loginRequest.Username, loginRequest.Password); !ok {
		respondWithError(responseWriter, "Wrong Credentials.", http.StatusUnauthorized)
		return
	}
	respondWithSuccess(responseWriter, "Logged.", http.StatusOK)
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
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	jsonResponse, _ := json.Marshal(i)
	responseWriter.WriteHeader(status)
	_, err := responseWriter.Write(jsonResponse)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
	}
}
