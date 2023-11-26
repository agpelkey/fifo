package main
/*
import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/agpelkey/food"
	"github.com/julienschmidt/httprouter"
)

// POST
func (app *application) handleCreateDryGoodsItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	newDryGoodsItem := &food.DryGoods{
		Item: jsonInput.Item,	
		Unit: jsonInput.Unit,
		Quantity: jsonInput.Quantity,
		Purchase_date: time.Now(),
	}

	err = newDryGoodsItem.ValidateDryGoods()
	if err != nil {
		app.failedValidationResponse(w, r, err)
		return
	}

	err = app.DryGoodStore.CreateNewDryGood(*newDryGoodsItem)
	if err != nil {
		app.serverErrorResponse(w, r, err)	
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/fifo/dairy/%s", newDryGoodsItem.Item))

	_ = writeJSON(w, http.StatusOK, envelope{"item": newDryGoodsItem}, headers)
}

// GET
func (app *application) handleGetDryGoodsItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	urlQueryName := queryUrlParam(ps)		

	dairyItem, err := app.DryGoodStore.GetDryGoodFromDB(urlQueryName)
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
func (app *application) handleUpdateDryGoods(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    itemResult, err := app.DryGoodStore.GetDryGoodFromDB(queryItemName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrDryGoodItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }
    
    var payload food.DryGoodItemUpdate

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

    err = app.DryGoodStore.UpdateDryGoodItem(itemResult)
    if err != nil {
        app.errorResponse(w, r, http.StatusConflict, fmt.Errorf("there was a problem editing the item"))
    }

    _ = writeJSON(w, http.StatusOK, envelope{"message": "item successfully updated"}, nil)
    
}

// DELETE
func (app *application) handleDeleteDryGoods(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    err := app.DryGoodStore.DeleteDryGoodItem(queryItemName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrDryGoodItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    err = writeJSON(w, http.StatusOK, envelope{"message": "item succesfully deleted"}, nil)
}
*/
