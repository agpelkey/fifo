package food

import (
	"errors"
	"time"
)

var (

	errDryGoodNameRequired = errors.New("Item required")
	errDryGoodUnitRequired = errors.New("Item units required")
	errDryGoodQuantityRequired = errors.New("Item quantity required")
	errDryGoodPurchaseDateRequired = errors.New("Purchase date required")

    	ErrDryGoodItemNotFound = errors.New("item could not be found")

)

type DryGoods struct {
	Item string `json:"item"`
	Unit string `json:"unit"`
	Quantity float32`json:"quantity"`
	Purchase_date time.Time `json:"purchase_date"`
}

type DryGoodService interface {
	CreateNewDryGood(item DryGoods) error
	GetDryGoodFromDB(item string) (*DryGoods, error)
	UpdateDryGoodItem(item *DryGoods) error
	DeleteDryGoodItem(item string) error
}


type DryGoodItemUpdate struct {
	Item *string `json:"item"`
	Unit *string `json:"unit"`
	Quantity *float32 `json:"quantity"`
}

func (d DryGoods) ValidateDryGoods() error {
	switch {
	case d.Item == "":
		return errDryGoodNameRequired
	case d.Unit == "":
		return errDryGoodUnitRequired
	case d.Quantity == 0:
		return errDryGoodQuantityRequired
	default:
		return nil
	}
}
