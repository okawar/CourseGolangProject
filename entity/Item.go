package entity

import "errors"

type Item struct {
	ItemId      uint32  `json:"item_id"`
	Article     uint32  `json:"article"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Creator     string  `json:"creator"`
	Description string  `json:"description"`
}

// Validate проверяет корректность полей Item.
// Возвращает срез ошибок, если какое-либо из полей недопустимо.
func (item *Item) Validate() []error {
	var errs []error

	if item.ItemId == 0 {
		errs = append(errs, errors.New("ItemId must be greater than 0"))
	}

	if item.Name == "" {
		errs = append(errs, errors.New("NAME CAN'T BE EMPTY"))
	}

	if item.Price <= 0 {
		errs = append(errs, errors.New("PRICE MUST BE > 0"))
	}

	return errs
}
