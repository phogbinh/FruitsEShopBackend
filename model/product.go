package model

type Product struct {
	Name         string `json:"Pname"`
	StaffName    string `json:"StaffUserName"`
	Description  string `json:"Description"`
	Category     string `json:"Category"`
	Source       string `json:"Source"`
	Price        int    `json:"Price"`
	Inventory    int    `json:"Inventory"`
	Quantity     int    `json:"SoldQuantity"`
	SaleDate     string `json:"OnSaleDate"`
}
