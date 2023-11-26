package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) handleGetItemFromFridge(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
     
    queryName := ps.ByName("item")

    result, err := app.ItemStore.GetItemFromFridge(queryName)
    if err != nil {
        app.notFoundResponse(w, r)
        //app.serverErrorResponse(w, r, err)
        return
    }

    err = writeJSON(w, http.StatusOK, envelope{"item": result}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }

}
