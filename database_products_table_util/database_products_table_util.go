package database_products_table_util

import (
	"database/sql"

	"backend/database_util"
	. "backend/model"
)

const (
	TableName                                = "product"
	IdColumnName                             = "ProductId"
	SpecialEventDiscountPolicyCodeColumnName = "SpecialEventDiscountPolicyCode"
)

const (
	queryUpdateProductSpecialEventDiscountPolicyCode = "UPDATE " + TableName + " SET " + SpecialEventDiscountPolicyCodeColumnName + " = ? WHERE " + IdColumnName + " = ?"
)

// UpdateProductSpecialEventDiscountPolicyCode updates the given product's special event discount policy code to the database `products` table.
func UpdateProductSpecialEventDiscountPolicyCode(productId string, productNewSpecialEventDiscountPolicyCode string, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryUpdateProductSpecialEventDiscountPolicyCode, productNewSpecialEventDiscountPolicyCode, productId)
}
