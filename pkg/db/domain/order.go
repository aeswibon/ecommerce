package domain

type Order struct {
	ID          string  `json:"id"`
	CustomerID  string  `json:"customer_id"`
	ProductID   string  `json:"product_id"`
	Quantity    int     `json:"quantity"`
	TotalAmount float64 `json:"total_amount"`
	Address     Address `json:"address"`
}
