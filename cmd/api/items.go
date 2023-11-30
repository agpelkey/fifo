package main

import (
	"fmt"
	"net/http"

	"github.com/agpelkey/food"
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

func (app *application) handleInsertItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var input struct {
        Name string `json:"name"`
        Type string `json:"type"`
        Unit string `json:"unit"`
        Quantity float64 `json:"quantity"`
    }

    err := readJSON(w, r, &input)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    newItem := &food.Items {
        Name: input.Name,
        Type: input.Type,
        Unit: input.Unit,
        Quantity: input.Quantity,
    }

    err = newItem.ValidateItemCreate()
    if err != nil {
        app.failedValidationResponse(w, r, err)
        return
    }

    err = app.ItemStore.InsertIntoFridge(*newItem)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    headers := make(http.Header)
    headers.Set("Location", fmt.Sprintf("/v1/fifo/%s", newItem.Name))

    _ = writeJSON(w, http.StatusOK, envelope{"item": newItem}, headers)

}
