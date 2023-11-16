package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/agpelkey/food"
	"github.com/julienschmidt/httprouter"
)

// POST
func (app *application) handleCreateDairyItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jsonInput struct {
		Item string `json:"item"`
		Unit string `json:"unit"`
		Quantity float32`json:"quantity"`
	}

	err := readJSON(w, r, &jsonInput)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	newDairyItem := &food.Dairy{
		Item: jsonInput.Item,	
		Unit: jsonInput.Unit,
		Quantity: jsonInput.Quantity,
		Purchase_date: time.Now(),
	}

	err = newDairyItem.ValidateDairyCreate()
	if err != nil {
		app.failedValidationResponse(w, r, err)
		return
	}

	err = app.DairyStore.CreateNewDairy(*newDairyItem)
	if err != nil {
		app.serverErrorResponse(w, r, err)	
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/fifo/dairy/%s", newDairyItem.Item))

	_ = writeJSON(w, http.StatusOK, envelope{"item": newDairyItem}, headers)
}

// GET
func (app *application) handleGetDairyItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	urlQueryName := queryUrlParam(ps)		

	dairyItem, err := app.DairyStore.GetDairyFromDB(urlQueryName)
	if err != nil {
		app.notFoundResponse(w, r)
		return 
	}

	err = writeJSON(w, http.StatusOK, envelope{"item": dairyItem}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}


// PATCH
func (app *application) handleUpdateDairy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    itemResult, err := app.DairyStore.GetDairyFromDB(queryItemName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrDairyItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }
    
    var payload food.DairyItemUpdate

    err = readJSON(w, r, &payload)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    if payload.Item != nil {
        itemResult.Item = *payload.Item
    }
    if payload.Unit!= nil {
        itemResult.Unit= *payload.Unit
    }
    if payload.Quantity != nil {
        itemResult.Quantity = *payload.Quantity
    }

    err = app.DairyStore.UpdateDairyItem(itemResult)
    if err != nil {
        app.errorResponse(w, r, http.StatusConflict, fmt.Errorf("there was a problem editing the item"))
    }

    _ = writeJSON(w, http.StatusOK, envelope{"message": "item successfully updated"}, nil)
    
}

// DELETE
func (app *application) handleDeleteDairy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    err := app.DairyStore.DeleteDairyItem(queryItemName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrDairyItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    err = writeJSON(w, http.StatusOK, envelope{"message": "item succesfully deleted"}, nil)
}
