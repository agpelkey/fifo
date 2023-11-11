package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// function to handle API routes
func (app *application) routes() http.Handler {
    r := httprouter.New()

    r.HandlerFunc(http.MethodPost, "/v1/fifo/protein", app.NewProteinRequest)

    return r
}
