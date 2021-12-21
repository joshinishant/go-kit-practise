package entity

type Order struct {
	Id           string      `json:"id,omitempty"`
	CustmerId    string      `json:"customer_id"`
	Status       string      `json:"status"`
	CreatedOn    int64       `json:"created_on,omitempty"`
	RestaurantId string      `json:"restaurant_id"`
	OrderItems   []OrderItem `json:"order_items,omitempty"`
}

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	Name        string  `json:"name"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}
