package food

import "errors"


var (

	errItemNameRequired = errors.New("Item required")
	errUnitRequired = errors.New("Item units required")
	errQuantityRequired = errors.New("Item quantity required")
	errPurchaseDateRequired = errors.New("Purchase date required")
	errTypeRequired = errors.New("Item type required")

    	ErrItemNotFound = errors.New("item not found, please enter type and unit")

)

type Items struct {
	Item_id int `json:"item_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Unit string `json:"unit"`
    Quantity float64 `json:"quantity"`
}

type ItemStore interface {
	GetItemFromFridge(name string) (Items, error)
    InsertIntoFridge(item Items) error
    UpdateFridgeItem(item Items) error
    DeleteItemFromFridge(item string) error
}

type ItemUpdate struct {
	Name *string `json:"name"`
	Type *string `json:"type"`
	Unit *string `json:"unit"`
	Quantity *float64`json:"quantity"`
    Item_id *int `json:"item_id"`
}

func (i Items) ValidateItemCreate() error {
	switch {
	case i.Name == "":
		return errItemNameRequired
	case i.Type == "":
		return errTypeRequired
	case i.Unit == "":
		return errUnitRequired
    case i.Quantity == 0:
        return errQuantityRequired
	default:
		return nil
	}
}
