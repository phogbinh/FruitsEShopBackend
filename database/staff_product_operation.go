package database

import (
	"database/sql"
	"encoding/json"
	"model/product.go"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var serialNumber int = 1000

// get new product id (no repeat serial number)
func GetNewProductId() (pid string){
	pid = strconv.Itoa(serialNumber)
	serialNumber++
	return pid
}

// Staff add product to database 
func AddProduct(info *Product, databasePtr *sql.DB) (code int) {
	_, addError := databasePtr.Exec("INSERT	INTO product (p_id, s_username, description, p_name, category, source, price, inventory, sold_quantity, onsale_date)\n"+
		"	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
		 GetNewProductId(), info.StaffName, info.Description, info.Name, info.Category, info.Source, info.Price, info.Inventory, info.Quantity, info.SaleDate)

	if addError != nil {
		fmt.Println(addError)
		statusCode = 417
	} else {
		statusCode = 200
	}
	return statusCode
}

func DeleteProduct(productID int, databasePtr *sql.DB) (code int){
	_, deleteError := databasePtr.Exec("DELETE FROM product\n"+
		"WHERE 	p_id = ?;" , productID)

	if deleteError != nil {
		fmt.Println(deleteError)
		statusCode = 417
	} else {
		statusCode = 200
	}
	return statusCode
}

// Modify order item quantity
func ModifyProduct(productID int, info *Product, databasePtr *sql.DB) (code int) {
	_, modifyError := databasePtr.Exec("UPDATE product\n"+
		"	SET	s_username = ?\n"+
		"	,	description = ?\n"+
		"	,	p_name = ?\n"+
		"	,	category = ?\n"+
		"	,	source = ?\n"+
		"	,	price = ?\n"+
		"	,	inventory = ?\n"+
		"	,	quantity = ?\n"+
		"WHERE	p_id = ?;" , info.StaffName, info.Description, info.Name, info.Category, info.Source, info.Price, info.Inventory, info.Quantity, productID)

	if modifyError != nil {
		fmt.Println(modifyError)
		statusCode = 403
	} else {
		statusCode = 200
	}

	return statusCode
}