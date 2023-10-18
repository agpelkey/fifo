package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// function to handle API routes
func (app *application) routes() http.Handler {
    r := httprouter.New()


    return r
}
