package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/agpelkey/food"
	"github.com/julienschmidt/httprouter"
)

func (app *application) handleHTMX(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("/fifo-frontend/index.html"))
    tmpl.Execute(w, nil)
}

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

func (app *application) handleGetItemByID(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r) 
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    item, err := app.FoodStore.GetItemByID(id)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrFridgeItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    _ = writeJSON(w, http.StatusOK, envelope{"item": item}, nil)

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

func (app *application) handleUpdateFridgeQuantity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    queryName := ps.ByName("item")

    fridgeItem, err := app.ItemStore.GetItemFromFridge(queryName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    var payload food.ItemUpdate

    err = readJSON(w, r, &payload)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    if payload.Item_id != nil {
        fridgeItem.Item_id = *payload.Item_id
    }

    if payload.Name != nil {
        fridgeItem.Name = *payload.Name
    }

    if payload.Type != nil {
        fridgeItem.Type = *payload.Type
    }

    if payload.Unit != nil {
        fridgeItem.Unit = *payload.Unit
    }

    if payload.Quantity != nil {
        fridgeItem.Quantity = *payload.Quantity
    }

    err = app.ItemStore.UpdateFridgeItem(fridgeItem)
    if err != nil {
        app.errorResponse(w, r, http.StatusConflict, fmt.Errorf("There was a problem editing the item"))
    }

    _ = writeJSON(w, http.StatusOK, envelope{"message:": "item was succesfully updated"}, nil)
}

func (app *application) handleDeleteItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    param := ps.ByName("item")

    err := app.ItemStore.DeleteItemFromFridge(param)

    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    _ = writeJSON(w, http.StatusOK, envelope{"item deleted": ""}, nil)
}


func (app *application) readIDParam(r *http.Request) (int64, error) {
    params := httprouter.ParamsFromContext(r.Context())

    id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
    if err != nil {
        return 0, errors.New("invalid id parameter")
    }

    return id, nil
}







