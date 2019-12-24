package database

import (
	. "backend/model"
	"database/sql"
	"encoding/json"
	"log"
)

/// GetOrder is a function which handle all order
func GetOrder(user User, databasePtr *sql.DB) (code int, jsonData string) {
	rows, getError := databasePtr.Query("SELECT " + ProductNameColumnName + ", " + ProductPriceColumnName + ", " + TradeDateTimeColumnName +
		" From " + "( " + "SELECT " + ProductIdColumnName + ", " + CustomerOwnCartCustomerUserNameColumnName + ", " + TradeDateTimeColumnName +
		" FROM " + CustomerOwnCartTableName + " as c join " + TradeTableName + " as t" +
		"	WHERE " + "c.CartId = t.CartId and c.CustomerUserName = \"jamfly\") as temp join " + ProductTableName + " where temp.ProductId = product.ProductId;")

	if getError != nil {
		log.Panicln(getError)
		return 500, "message: failed on get order"
	}

	defer rows.Close()

	code = 200

	columns, err := rows.Columns()

	if err != nil {
		return 500, "failed on get order data"
	}

	tableData := make([]map[string]interface{}, 0)
	// TODO: need to transfer to correct data
	appendRowsDataIntoTableData(rows, tableData, columns)

	json, err := json.Marshal(tableData)

	if err != nil {
		return 500, "failed on get order data"
	}

	jsonData = string(json)

	return 200, jsonData
}
