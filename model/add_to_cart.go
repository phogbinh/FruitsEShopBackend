package model

type AddToCart struct {
	ProductID int `json:"p_id"`
	CartID    int `json:"c_id"`
	Quantity  int `json:"quantity"`
}
