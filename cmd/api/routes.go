package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// function to handle API routes
func (app *application) routes() http.Handler {
    r := httprouter.New()

    r.POST("/v1/fifo/protein", app.handleNewProtein)
    r.GET("/v1/fifo/protein/:item", app.handleGetProtein)
    r.PATCH("/v1/fifo/protein/:item", app.handleUpdateProtein)
    r.DELETE("/v1/fifo/protein/:item", app.handleDeleteProtein)

    r.POST("/v1/fifo/dairy", app.handleCreateDairyItem)
    r.GET("/v1/fifo/dairy/:item", app.handleGetDairyItem)

    return r
}
