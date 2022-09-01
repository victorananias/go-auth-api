package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/victorananias/go-auth-api/responses"
)

type baseController struct{}

func (controller *baseController) decodeRequest(request *http.Request, object interface{}) error {
	return json.NewDecoder(request.Body).Decode(&object)
}

func (controller *baseController) respondWithError(responseWriter http.ResponseWriter, message string, status int) {
	response := responses.DefaultResponse{Message: message, Success: false}
	controller.respondWithJson(responseWriter, response, status)
}

func (controller *baseController) respondWithSuccess(responseWriter http.ResponseWriter, message string, status int) {
	response := responses.DefaultResponse{Message: message, Success: true}
	controller.respondWithJson(responseWriter, response, status)
}

func (controller *baseController) respondWithJson(responseWriter http.ResponseWriter, i interface{}, status int) {
	jsonResponse, _ := json.Marshal(i)
	responseWriter.WriteHeader(status)
	_, err := (responseWriter).Write(jsonResponse)
	if err != nil {
		http.Error((responseWriter), err.Error(), http.StatusBadRequest)
	}
}
