package model

type OrderItem struct {
	ProductID int `json:"ProductId"`
	CartID    int `json:"CartId"`
	Quantity  int `json:"Quantity"`
}
