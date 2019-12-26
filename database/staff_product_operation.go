package database

import (
	"database/sql"
	"encoding/json"
	. "backend/model"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

var serialNumber int = 1000

// get new product id (no repeat serial number)
func GetNewProductID() (pid string){
	pid = strconv.Itoa(serialNumber)
	serialNumber++
	return pid
}

// Staff add product to database 
func AddProduct(info *Product, databasePtr *sql.DB) (code int) {
	_, addError := databasePtr.Exec("INSERT	INTO "+ProductTableName+" ("+ProductIdColumnName+", "+ProductStaffUserNameColumnName+", "+ProductDescriptionColumnName+", "+ProductNameColumnName+", "+ProductCategoryColumnName+", "+ProductSourceColumnName+", "+ProductPriceColumnName+" , "+ProductInventoryColumnName+", "+ProductSoldQuantityColumnName+", "+ProductOnSaleDateColumnName+", "+productDiscountPolicyCodeColumnName+")\n" +
	" 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);" , GetNewProductID(), info.StaffName, info.Description, info.Name, info.Category, info.Source, info.Price, info.Inventory, info.Quantity, info.SaleDate, "000000000")

	if addError != nil {
		fmt.Println(addError)
		code = 417
	} else {
		code = 200
	}
	return code
}

// Delate Product
func DeleteProduct(productID int, databasePtr *sql.DB) (statusCode int){
	_, deleteError := databasePtr.Exec("DELETE FROM "+ProductTableName+"\n"+
		"WHERE 	"+ProductIdColumnName+" = ?;" , productID)

	if deleteError != nil {
		fmt.Println(deleteError)
		statusCode = 417
	} else {
		statusCode = 200
	}
	return statusCode
}

// Modify Product
func ModifyProduct(productID int, info *Product, databasePtr *sql.DB) (statusCode int) {
	_, modifyError := databasePtr.Exec("UPDATE "+ProductTableName+"\n"+
		"	SET	"+ProductStaffUserNameColumnName+" = ?\n"+
		"	,	"+ProductDescriptionColumnName+" = ?\n"+
		"	,	"+ProductNameColumnName+" = ?\n"+
		"	,	"+ProductCategoryColumnName+" = ?\n"+
		"	,	"+ProductSourceColumnName+" = ?\n"+
		"	,	"+ProductPriceColumnName+" = ?\n"+
		"	,	"+ProductInventoryColumnName+" = ?\n"+
		"	,	"+ProductSoldQuantityColumnName+" = ?\n"+
		"WHERE	"+ProductIdColumnName+" = ?;" , info.StaffName, info.Description, info.Name, info.Category, info.Source, info.Price, info.Inventory, info.Quantity, productID)

	if modifyError != nil {
		fmt.Println(modifyError)
		statusCode = 403
	} else {
		statusCode = 200
	}

	return statusCode
}

// Query Product By ProductName or StaffName
func QueryProduct(ProductName string, staffName string, databasePtr *sql.DB) (code int, jsonData string) {
	rows, queryError := databasePtr.Query("SELECT	"+ProductNameColumnName+",  "+ProductPriceColumnName+"\n" +
		"	FROM	"+ProductTableName+"\n"+
		"	WHERE	"+ProductNameColumnName+" = ?\n"+
		" 	OR "+ProductStaffUserNameColumnName+" = ?" , ProductName , staffName)

	if queryError != nil {
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
	appendRowsDataIntoTableData(rows, &tableData, columns)
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