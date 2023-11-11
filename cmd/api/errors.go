package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	responseToUsers := envelope{"error:": message}

	err := writeJSON(w, status, responseToUsers, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)

	messageToUsers := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, messageToUsers)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	messageToUsers := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, messageToUsers)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	messageToUsers := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, messageToUsers)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, err)
} 
