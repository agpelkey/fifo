package food

import (
	"errors"
	"time"
)


var (

	errFruitNameRequired = errors.New("Item required")
	errFruitUnitRequired = errors.New("Item units required")
	errFruitQuantityRequired = errors.New("Item quantity required")
	errFruitPurchaseDateRequired = errors.New("Purchase date required")

    	ErrFruitItemNotFound = errors.New("item could not be found")

)

type Fruit struct {
	Item string `json:"item"`
	Unit string `json:"unit"`
	Quantity float32 `json:"quantity"`
	Purchase_date time.Time `json:"purchase_dat"`
}

type FruitService interface {
	CreateNewFruit(item Fruit) error
	GetFruitFromDB(item string) (*Fruit, error)
	UpdateFruitItem(item *Fruit) error
	DeleteFruitItem(item string) error
}


type FruitItemUpdate struct {
	Item *string `json:"item"`
	Unit *string `json:"unit"`
	Quantity *float32 `json:"quantity"`
} 

func (f Fruit) ValidateFruitCreate() error {
	switch {
	case f.Item == "":
		return errFruitNameRequired
	case f.Unit == "":
		return errFruitUnitRequired
	case f.Quantity == 0:
		return errFruitQuantityRequired
	default:
		return nil
	}
}
