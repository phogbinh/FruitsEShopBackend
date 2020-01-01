package database

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

func GetStaffOrder(staffUserName string, databasePtr *sql.DB) (code int, jsonData string) {
	rows, err := databasePtr.Query("SELECT "+CustomerOwnCartCustomerUserNameColumnName+", "+
		ProductNameColumnName+", "+TradeProductQuantityColumnName+", "+TradeDateTimeColumnName+"\n"+
		"	FROM "+TradeTableName+" INNER JOIN "+ProductTableName+" ON "+
		TradeTableName+"."+TradeProductIdColumnName+" = "+ProductTableName+"."+ProductIdColumnName+"\n"+
		"	INNER JOIN "+CustomerOwnCartTableName+" ON "+
		TradeTableName+"."+TradeCartIdColumnName+" = "+CustomerOwnCartTableName+"."+CustomerOwnCartCartIdColumnName+"\n"+
		"	WHERE "+ProductTableName+"."+ProductStaffUserNameColumnName+" = ?"+
		"	ORDER BY "+TradeTableName+"."+TradeDateTimeColumnName+";", staffUserName)

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
	appendRowsDataIntoTableData(rows, &tableData, columns)
	json, err := json.Marshal(tableData)

	if err != nil {
		code, jsonData = setFailureDataForGetOrderItemsInCart()
		return code, jsonData
	}

	jsonData = string(json)

	return 200, jsonData
}
