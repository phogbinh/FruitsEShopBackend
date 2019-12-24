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

// Query Product By ProductName or StaffName
func QueryProduct(ProductName string, staffName string, databasePtr *sql.DB) (code int) {
	rows, queryError := databasePtr.Query("SELECT	p_name,  price\n" +
		"	FROM	product\n"+
		"	WHERE	p_name = ?\n"+
		" 	OR s_username = ?" , ProductName , staffName)

	if err != nil {
		code, jsonData = setFailureDataForQueryProduct();
		return code, jsonData
	}
	
	defer rows.Close()
	
	code = 200

	columns, err := rows.Columns()
	
	if err != nil {
		code, jsonData = setFailureDataForQueryProduct();
		return code, jsonData
	}
	
	tableData := make([]map[string]interface{}, 0)
	appendRowsDataIntoTableData(rows, tableData, columns)
	json, err := json.Marshal(tableData)

	if err != nil {
		code, jsonData = setFailureDataForQueryProduct();
		return code, jsonData
	}

	jsonData = string(json)
	return 200, jsonData
}

// set function "QueryProduct" failure data
func setFailureDataForQueryProduct() (code int, jsonData string) {
	code = 403
	jsonData = ""
	return code, jsonData
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