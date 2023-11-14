package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/agpelkey/food"
	"github.com/julienschmidt/httprouter"
)

func (app application) NewProteinRequest(w http.ResponseWriter, r *http.Request) {
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
func (app *application) GetProteinItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    //queryItemName := r.URL.Query().Get("item") 

    queryItemName := ps.ByName("item")

    fmt.Println(queryItemName)

    result, err := app.ProteinStore.GetProteinFromDB(queryItemName)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }

    _ = writeJSON(w, http.StatusOK, envelope{"item": result}, nil)

}


// Update an item


// Delete an item
