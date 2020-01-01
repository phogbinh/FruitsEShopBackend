package database

import (
	. "backend/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/// TransactionFromCart do the transaction while buying
func TransactionFromCart(addToCart *OrderItem, databasePtr *sql.DB) (code int) {
	_, err := databasePtr.Exec("INSERT INTO "+TradeTableName+"("+TradeCartIdColumnName+", "+TradeProductIdColumnName+", "+TradeProductQuantityColumnName+", "+TradeDateTimeColumnName+")\n"+
		"SELECT CartId, ProductId, Quantity, Now()\n"+
		"FROM "+OrderItemTableName+"\n"+
		"WHERE "+OrderItemCartIdColumnName+"= ?;\n", addToCart.CartID)
	if err != nil {
		return 500
	}

	_, delete_all_item_err := databasePtr.Exec("UPDATE " + ProductTableName + ", " + OrderItemTableName +
		" SET  " + ProductTableName + "." + ProductSoldQuantityColumnName + " = " +
		ProductTableName + "." + ProductSoldQuantityColumnName + " + " + OrderItemTableName + "." + OrderItemQuantity +
		", " + ProductTableName + "." + ProductInventoryColumnName + " = " +
		ProductTableName + "." + ProductInventoryColumnName + " - " + OrderItemTableName + "." + OrderItemQuantity +
		" WHERE " + ProductTableName + "." + ProductIdColumnName + " = " + OrderItemTableName + "." + OrderItemProductIdColumnName + ";")

	if delete_all_item_err != nil {
		return 500
	}

	_, delete_quantity_err := databasePtr.Exec("DELETE FROM "+OrderItemTableName+" WHERE CartId = ?;", addToCart.CartID)
	if delete_quantity_err != nil {
		return 500
	}

	return 200
}
