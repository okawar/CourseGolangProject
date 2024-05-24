package entity

// Order представляет сущность корзины
type Order struct {
	OrderID uint32  `json:"order_id"`
	UserID  uint32  `json:"user_id"`
	Amount  float32 `json:"amount"`
	ItemID  uint32  `json:"item_id"`
}
