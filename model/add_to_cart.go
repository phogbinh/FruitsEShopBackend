package model

type AddToCart struct {
	ProductID int `json:"p_id"`
	CartID    int `json:"c_id"`
	Quantity  int `json:"quantity"`
}

func (addToCart *AddToCart) SetProductID(productID int) error {
	addToCart.ProductID = productID
	return nil
}

func (addToCart *AddToCart) SetCartID(cartID int) error {
	addToCart.CartID = cartID
	return nil
}

func (addToCart *AddToCart) SetQuantity(quantity int) error {
	addToCart.Quantity = quantity
	return nil
}
