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

    ErrProteinItemNotFound = errors.New("item could not be found")
    ErrEditConflict = errors.New("item could not be updated")
)

// create custom protein data type
type Protein struct {
	Item string `json:"item"`
	Unit string `json:"unit"`
	Quantity float32 `json:"quantity"`
	Purchase_date time.Time `json:"purchase_date"`
}

type ProteinService interface {
   CreateNewProtein(item Protein) error
   GetProteinFromDB(name string) (*Protein, error)
   UpdateProteinItem(item *Protein) error
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
	default:
		return nil
	}
}

// validate for PATCH requests



