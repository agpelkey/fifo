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


/*
func (app application) handleNewProtein(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input struct {
		Item string `json:"item"`
		Unit string `json:"unit"`
		Quantity float32 `json:"quantity"`
		Purchase_date time.Time `json:"purchase_date"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	NewProteinItem := &food.Protein{
		Item: input.Item,
		Unit: input.Unit,
		Quantity: input.Quantity,
		Purchase_date: time.Now(),
	}

	err = NewProteinItem.ValidateProtein()
	if err != nil {
		app.failedValidationResponse(w, r, err)
		return
	}
	
	err = app.ProteinStore.CreateNewProtein(*NewProteinItem)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/fifo/protein/%s", input.Item))

	err = writeJSON(w, http.StatusOK, envelope{"protein item added": NewProteinItem}, headers)
}


// Get an item
func (app *application) handleGetProtein(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    result, err := app.ProteinStore.GetProteinFromDB(queryItemName)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, envelope{"error:": "item could not be found"}, nil)
        return
    }

    err = writeJSON(w, http.StatusOK, envelope{"item":result}, nil)
    if err != nil {
        app.notFoundResponse(w, r)
    }

}


// Update an item
func (app *application) handleUpdateProtein(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    itemResult, err := app.ProteinStore.GetProteinFromDB(queryItemName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrProteinItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }
    
    var payload food.ProteinUpdate

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

    err = app.ProteinStore.UpdateProteinItem(itemResult)
    if err != nil {
        app.errorResponse(w, r, http.StatusConflict, fmt.Errorf("there was a problem editing the item"))
    }

    _ = writeJSON(w, http.StatusOK, envelope{"message": "item successfully updated"}, nil)
    
}

// Delete an item
func (app *application) handleDeleteProtein(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryItemName := queryUrlParam(ps)

    err := app.ProteinStore.DeleteProteinItem(queryItemName)
    if err != nil {
        switch {
        case errors.Is(err, food.ErrProteinItemNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    err = writeJSON(w, http.StatusOK, envelope{"message": "item succesfully deleted"}, nil)
}

*/
