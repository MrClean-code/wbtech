package wbtech

type OrderDto struct {
	Order    Order    `json:"order" db:"order"`
	Delivery Delivery `json:"delivery" db:"deliver"`
	Payment  Payment  `json:"payment" db:"payment"`
	Items    []Item   `json:"items" db:"items"`
}
