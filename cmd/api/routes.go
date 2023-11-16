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
    r.PATCH("/v1/fifo/dairy/:item", app.handleUpdateDairy)
    r.DELETE("/v1/fifo/dairy/:item", app.handleDeleteDairy)

    r.POST("/v1/fifo/dry_goods", app.handleCreateDryGoodsItem)
    r.GET("/v1/fifo/dry_goods/:item", app.handleGetDryGoodsItem)
    r.PATCH("/v1/fifo/dry_goods/:item", app.handleUpdateDryGoods)
    r.DELETE("/v1/fifo/dry_goods/:item", app.handleDeleteDryGoods)

    r.POST("/v1/fifo/fruit", app.)
    r.GET("/v1/fifo/fruit/item", app.)
    r.PATCH("/v1/fifo/fruit/:item", app.)
    r.DELETE("/v1/fifo/fruit/:item", app.)


    return r
}
