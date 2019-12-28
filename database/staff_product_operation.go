package database

import (
	"database/sql"
	"encoding/json"
	. "backend/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var serialNumber int = 1000

// get new product id (no repeat serial number)
func GetNewProductID(databasePtr *sql.DB) (pid int){
	var productID int
	err := databasePtr.QueryRow("SELECT MAX("+ProductIdColumnName+")\n" +
	"	FROM	"+ProductTableName+"\n").Scan(&productID)
	if err != nil {
		fmt.Println(err)
	} else {
		pid = productID + 1
	}
	return pid
}

// Staff add product to database 
func AddProduct(info *Product, databasePtr *sql.DB) (code int) {
	_, addError := databasePtr.Exec("INSERT	INTO "+ProductTableName+" ("+ProductIdColumnName+", "+ProductStaffUserNameColumnName+", "+ProductDescriptionColumnName+", "+ProductNameColumnName+", "+ProductCategoryColumnName+", "+ProductSourceColumnName+", "+ProductPriceColumnName+" , "+ProductInventoryColumnName+", "+ProductSoldQuantityColumnName+", "+ProductOnSaleDateColumnName+", "+ProductImageSourceColumnName+")\n" +
	" 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);" , GetNewProductID(databasePtr), info.StaffName, info.Description, info.Name, info.Category, info.Source, info.Price, info.Inventory, info.Quantity, info.SaleDate, info.ImageSrc)

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
	fmt.Println(productID)
	fmt.Println(info)
	_, modifyError := databasePtr.Exec("UPDATE "+ProductTableName+"\n"+
		"	SET	"+ProductStaffUserNameColumnName+" = ?\n"+
		"	,	"+ProductDescriptionColumnName+" = ?\n"+
		"	,	"+ProductNameColumnName+" = ?\n"+
		"	,	"+ProductCategoryColumnName+" = ?\n"+
		"	,	"+ProductSourceColumnName+" = ?\n"+
		"	,	"+ProductPriceColumnName+" = ?\n"+
		"	,	"+ProductInventoryColumnName+" = ?\n"+
		"	,	"+ProductSoldQuantityColumnName+" = ?\n"+
		"	,	"+ProductImageSourceColumnName+" = ?\n"+
		"WHERE	"+ProductIdColumnName+" = ?;" , info.StaffName, info.Description, info.Name, info.Category, info.Source, info.Price, info.Inventory, info.Quantity, info.ImageSrc, productID)

	if modifyError != nil {
		fmt.Println(modifyError)
		statusCode = 403
	} else {
		statusCode = 200
	}

	return statusCode
}

// Query Product By ProductName or StaffName
func QueryProduct(ProductName string, StaffName string, ProductId int, databasePtr *sql.DB) (code int, jsonData string) {
	if ProductId != 0 {
		ProductName = "NULLITEM"
		StaffName = "NULLSTAFF"
	}else if StaffName != "" {
		ProductName = "NULLITEM"
		ProductId = 0
	}
	rows, queryError := databasePtr.Query("SELECT *\n" +
	"	FROM	"+ProductTableName+"\n"+
	"	WHERE	"+ProductNameColumnName+" like ?\n" +
	"	OR 		"+ProductStaffUserNameColumnName+" = ?\n" +
	"	OR 		"+ProductIdColumnName+" = ?" , "%" + ProductName + "%", StaffName, ProductId)

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