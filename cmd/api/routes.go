package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// function to handle API routes
func (app *application) routes() http.Handler {
    r := httprouter.New()

    r.HandlerFunc(http.MethodPost, "/v1/fifo/protein", app.NewProteinRequest)
    //r.HandlerFunc(http.MethodGet, "/v1/fifo/protein/:name", app.GetProteinItem)
    r.GET("/v1/fifo/protein/:item", app.GetProteinItem)

    return r
}
