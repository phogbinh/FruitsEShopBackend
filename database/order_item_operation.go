package database

import (
	. "backend/model"
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

// Add order product to cart
func AddOrderItemToCart(addToCart *AddToCart, databasePtr *sql.DB) (code int) {
	_, addError := databasePtr.Exec("INSERT	INTO add_to_cart (cart_id, p_id, quantity)\n"+
		"	VALUES (?, ?, ?);", addToCart.CartID, addToCart.ProductID, addToCart.Quantity)

	if addError != nil {
		code = 417
	} else {
		code = 200
	}
	return code
}

// Delete order item in cart
func DeleteOrderItemInCart(addToCart *AddToCart, databasePtr *sql.DB) (code int) {
	_, deleteError := databasePtr.Exec("DELETE FROM add_to_cart\n"+
		"	WHERE 	p_id = ?\n"+
		"	AND		cart_id = ?;", addToCart.ProductID, addToCart.CartID)

	if deleteError != nil {
		code = 417
	} else {
		code = 200
	}
	return code
}

// Get all order items in cart
func GetOrderItemsInCart(addToCart *AddToCart, databasePtr *sql.DB) (code int, jsonData string) {
	rows, err := databasePtr.Query("SELECT	p_name, category, description, source, price, inventory\n"+
		"	FROM	product, add_to_cart\n"+
		"	WHERE	cart_id = ?\n"+
		"	AND		add_to_cart.p_id = product.p_id", addToCart.CartID)

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	defer rows.Close()

	code = 200

	columns, err := rows.Columns()

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	tableData := make([]map[string]interface{}, 0)
	appendRowsDataIntoTableData(rows, tableData, columns)
	json, err := json.Marshal(tableData)

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	jsonData = string(json)

	return 200, jsonData
}

func appendRowsDataIntoTableData(rows *sql.Rows, tableData []map[string]interface{}, columns []string) {
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		putScanValuesIntoValues(valuePtrs, values, rows, count)
		entry := make(map[string]interface{})
		putColumnsDataIntoEntry(entry, columns, values)
		tableData = append(tableData, entry)
	}
}

func putScanValuesIntoValues(valuePtrs []interface{}, values []interface{}, rows *sql.Rows, count int) {
	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	rows.Scan(valuePtrs...)
}

func putColumnsDataIntoEntry(entry map[string]interface{}, columns []string, values []interface{}) {
	for i, col := range columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		entry[col] = v
	}
}

// set function "GetOrderItemsInCart" failure data
func setFailureDataForGetOrderItemsInCart() (code int, jsonData string) {
	code = 403
	jsonData = ""
	return code, jsonData
}

// Modify order item quantity
func ModifyOrderItemQuantity(addToCart *AddToCart, databasePtr *sql.DB) (code int) {
	_, modifyError := databasePtr.Exec("UPDATE add_to_cart\n"+
		"	SET		quantity = ?\n"+
		"	WHERE	p_id = ?\n"+
		"	AND		cart_id = ?;", addToCart.Quantity, addToCart.ProductID, addToCart.CartID)

	if modifyError != nil {
		code = 403
	} else {
		code = 200
	}

	return code
}
