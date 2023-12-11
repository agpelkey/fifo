package food

import (
	"errors"
	"time"
)

var (

	errFridgeQuantityRequired = errors.New("Item quantity required")
	errFridgePurchaseDateRequired = errors.New("Purchase date required")

    ErrFridgeItemNotFound = errors.New("item not found, please enter type and unit")

)

type Fridge struct {
    Item_id int `json:"item_id"`
    Quantity float32 `json:"quantity"`
    Purchase_date time.Time `json:"purchase_date"`
}


type FridgeStore interface {
    //UpdateFridgeItem(fridge Fridge) error
    GetItemByID(id int64) (Fridge, error)
}


type FridgeUpdate struct {
    Quantity *float32 `json:"quantity"`
}
