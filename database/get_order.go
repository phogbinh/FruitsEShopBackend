package database

import (
	"database/sql"
	"encoding/json"
	"log"
)

/// GetOrder is a function which handle all order
func GetOrder(cartId int, databasePtr *sql.DB) (code int, jsonData string) {
	rows, getError := databasePtr.Query("SELECT "+ProductNameColumnName+", "+ProductPriceColumnName+", "+TradeDateTimeColumnName+
		" From "+"( "+"SELECT "+ProductIdColumnName+", "+TradeDateTimeColumnName+
		" FROM "+TradeTableName+" where CartId = ?) as t join "+ProductTableName+" as p"+
		"	ON "+"t.ProductId = p.ProductId "+
		" ORDER BY "+"t.DateTime", cartId)

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
	appendRowsDataIntoTableData(rows, &tableData, columns)
	json, err := json.Marshal(tableData)

	if err != nil {
		return 500, "failed on get order data"
	}

	jsonData = string(json)

	return 200, jsonData
}
