package database_discount_policy_types_table_util

import (
	"database/sql"
	"errors"

	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	"backend/database_util"
	. "backend/model"
	"backend/util"
)

const (
	tableName      = discountPoliciesTablesConst.DiscountPolicyTypesTableName
	nameColumnName = discountPoliciesTablesConst.DiscountPolicyTypesNameColumnName
)

const (
	queryCreateTable = "CREATE TABLE IF NOT EXISTS " + tableName + util.EndOfLine +
		"(" + util.EndOfLine +
		nameColumnName + "	VARCHAR(255)	NOT NULL," + util.EndOfLine +
		"PRIMARY KEY(" + nameColumnName + ")" + util.EndOfLine +
		")"
	queryInsertTypeIfNotExists = "INSERT INTO " + tableName + util.EndOfLine +
		"VALUES						(?)" + util.EndOfLine +
		"ON DUPLICATE KEY UPDATE	" + nameColumnName + " = ?"
)

const (
	TypeShipping     = "Shipping"
	TypeSeasonings   = "Seasonings"
	TypeSpecialEvent = "Special Event"
)

// Initialize creates table `discount_policy_types` and inserts initial types if not exist.
func Initialize(databasePtr *sql.DB) error {
	createTableError := database_util.CreateTableIfNotExists(databasePtr, queryCreateTable)
	if createTableError != nil {
		return createTableError
	}
	insertError := insertInitialTypesIfNotExist(databasePtr)
	if insertError != nil {
		return insertError
	}
	return nil
}

func insertInitialTypesIfNotExist(databasePtr *sql.DB) error {
	var insertStatus Status
	insertStatus = insertTypeIfNotExists(TypeShipping, databasePtr)
	if !util.IsStatusOK(insertStatus) {
		return errors.New(insertStatus.ErrorMessage)
	}
	insertStatus = insertTypeIfNotExists(TypeSeasonings, databasePtr)
	if !util.IsStatusOK(insertStatus) {
		return errors.New(insertStatus.ErrorMessage)
	}
	insertStatus = insertTypeIfNotExists(TypeSpecialEvent, databasePtr)
	if !util.IsStatusOK(insertStatus) {
		return errors.New(insertStatus.ErrorMessage)
	}
	return nil
}

func insertTypeIfNotExists(typeName string, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryInsertTypeIfNotExists, typeName, typeName)
}

// IsValidDiscountPolicyType returns true if the given discount policy type is valid.
func IsValidDiscountPolicyType(discountPolicyType string) bool {
	return discountPolicyType == TypeShipping || discountPolicyType == TypeSeasonings || discountPolicyType == TypeSpecialEvent
}
