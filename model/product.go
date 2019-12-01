package model

type Product struct {
	Name         string `json:"p_name"`
	StaffName    string `json:"staffName"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Source       string `json:"source"`
	Price        int    `json:"price"`
	Inventory    int    `json:"inventory"`
	Quantity     int    `json:"quantity"`
	SaleDate     string `json:"saleDate"`
}
