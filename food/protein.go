package food

import (
	"errors"
	"time"
)


var (
	errItemNameRequired = errors.New("Item required")
	errUnitRequired = errors.New("Item units required")
	errQuantityRequired = errors.New("Item quantity required")
	errPurchaseDateRequired = errors.New("Purchase date required")
)

// create custom protein data type
type Protein struct {
	Item string `json:"item"`
	Unit string `json:"unit"`
	Quantity float32 `json:"quantity"`
	Purchase_date time.Time `json:"purchase_date"`
}

// protein db interface
type ProteinService interface {
   // DB logic goes here 
}


// Update protein item
type ProteinUpdate struct {
	Item *string `json:"item"`
	Unit *string `json:"unit"`
	Quantity *float32 `json:"quantity"`
	Purchase_date *time.Time `json:"purchase_date"`
}


// Validate POST requests
func (p Protein) ValidateProtein() error {
	switch {
	case p.Item == "":
		return errItemNameRequired
	case p.Unit == "":
		return errUnitRequired
	case p.Quantity == 0:
		return errQuantityRequired
	case p.Purchase_date != time.Now():
		return errPurchaseDateRequired // definitely will need to test this
	default:
		return nil
	}
}

// validate for PATCH requests



