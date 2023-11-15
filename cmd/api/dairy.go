package main

import (
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


// DELETE
