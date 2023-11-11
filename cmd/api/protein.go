package main

import (
	"fmt"
	"net/http"

	"github.com/agpelkey/food"
)

func (app application) NewProteinRequest(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Item string `json:"item"`
		Unit string `json:"unit"`
		Quantity float32 `json:"quantity"`
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

	err = writeJSON(w, http.StatusOK, envelope{"protein item added": input}, headers)
}
