package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Add order product to cart
func AddOrderItemToCart(productID int, cartID int, quantity int, databasePtr *sql.DB) (code int) {
	_, addError := databasePtr.Exec("INSERT	INTO add_to_cart (cart_id, p_id, quantity)\n"+
		"	VALUES (?, ?, ?);", cartID, productID, quantity)

	if addError != nil {
		fmt.Println("cartID" + strconv.Itoa(cartID))
		fmt.Println("productID" + strconv.Itoa(productID))
		fmt.Println(addError)
		code = 417
	} else {
		code = 200
	}
	return code
}

// Delete order item in cart
func DeleteOrderItemInCart(productID int, cartID int, databasePtr *sql.DB) (code int) {
	_, deleteError := databasePtr.Exec("DELETE FROM add_to_cart\n"+
		"	WHERE 	p_id = ?\n"+
		"	AND		cart_id = ?;", productID, cartID)

	if deleteError != nil {
		code = 417
	} else {
		code = 200
	}
	return code
}

// Get all order items in cart
func GetOrderItemsInCart(cartID int, databasePtr *sql.DB) (code int, jsonData string) {
	rows, err := databasePtr.Query("SELECT	p_name, category, description, source, price, inventory\n"+
		"	FROM	product, add_to_cart\n"+
		"	WHERE	cart_id = ?\n"+
		"	AND		add_to_cart.p_id = product.p_id", cartID)

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	defer rows.Close()

	columns, err := rows.Columns()

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	code = 200

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
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
		tableData = append(tableData, entry)
	}

	json, err := json.Marshal(tableData)

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	jsonData = string(json)

	return 200, jsonData
}

// set function "GetOrderItemsInCart" failure data
func setFailureDataForGetOrderItemsInCart() (code int, jsonData string) {
	code = 403
	jsonData = ""
	return code, jsonData
}

// Modify order item quantity
func ModifyOrderItemQuantity(productID int, cartID int, quantity int, databasePtr *sql.DB) (code int) {
	_, modifyError := databasePtr.Exec("UPDATE add_to_cart\n"+
		"	SET		quantity = ?\n"+
		"	WHERE	p_id = ?\n"+
		"	AND		cart_id = ?;", quantity, productID, cartID)

	if modifyError != nil {
		code = 403
	} else {
		code = 200
	}

	return code
}
