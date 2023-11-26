package food

import (
	"errors"
	"time"
)


var (

	errDairyNameRequired = errors.New("Item required")
	errDairyUnitRequired = errors.New("Item units required")
	errDairyQuantityRequired = errors.New("Item quantity required")
	errDairyPurchaseDateRequired = errors.New("Purchase date required")

    	ErrDairyItemNotFound = errors.New("item could not be found")

)

type Dairy struct {
	Item string `json:"item"`
	Unit string `json:"unit"`
	Quantity float32 `json:"quantity"`
	Purchase_date time.Time `json:"purchase_date"`
}

type DairyService interface {
	CreateNewDairy(item Dairy) error
	GetDairyFromDB(item string) (*Dairy, error) 
	UpdateDairyItem(item *Dairy) error 
	DeleteDairyItem(item string) error
}


type DairyItemUpdate struct {
	Item *string `json:"item"`
	Unit *string `json:"unit"`
	Quantity *float32 `json:"quantity"`
}


func (d Dairy) ValidateDairyCreate() error {
	switch {
	case d.Item == "":
		return errDairyNameRequired
	case d.Unit == "":
		return errDairyUnitRequired
	case d.Quantity == 0:
		return errDairyQuantityRequired
	default:
		return nil
	}

}
