package database_discount_policies_table_util

import (
	"database/sql"

	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	usersTable "backend/database_users_table_util"
	"backend/database_util"
	"backend/util"
)

const (
	tableName               = discountPoliciesTablesConst.DiscountPoliciesTableName
	codeColumnName          = discountPoliciesTablesConst.DiscountPoliciesCodeColumnName
	nameColumnName          = discountPoliciesTablesConst.DiscountPoliciesNameColumnName
	descriptionColumnName   = discountPoliciesTablesConst.DiscountPoliciesDescriptionColumnName
	typeColumnName          = discountPoliciesTablesConst.DiscountPoliciesTypeColumnName
	staffUserNameColumnName = discountPoliciesTablesConst.DiscountPoliciesStaffUserNameColumnName
)

const (
	queryCreateTable = "CREATE TABLE IF NOT EXISTS " + tableName + util.EndOfLine +
		"(" + util.EndOfLine +
		codeColumnName + "			CHAR(9)			NOT NULL," + util.EndOfLine +
		nameColumnName + "			VARCHAR(255)	NOT NULL," + util.EndOfLine +
		descriptionColumnName + "			VARCHAR(255)	NOT NULL," + util.EndOfLine +
		typeColumnName + "			VARCHAR(255)	NOT NULL," + util.EndOfLine +
		staffUserNameColumnName + "	VARCHAR(30)		NOT NULL," + util.EndOfLine +
		"PRIMARY KEY(" + codeColumnName + ")," + util.EndOfLine +
		"FOREIGN KEY(" + typeColumnName + ") REFERENCES " + discountPoliciesTablesConst.DiscountPolicyTypesTableName + "(" + discountPoliciesTablesConst.DiscountPolicyTypesNameColumnName + ")" + util.EndOfLine +
		"	ON UPDATE CASCADE," + util.EndOfLine +
		"FOREIGN KEY(" + staffUserNameColumnName + ") REFERENCES " + usersTable.TableName + "(" + usersTable.UserNameColumnName + ")" + util.EndOfLine +
		"	ON DELETE CASCADE" + util.EndOfLine +
		")"
)

// CreateTableIfNotExists creates table `discount_policies`.
func CreateTableIfNotExists(databasePtr *sql.DB) error {
	return database_util.CreateTableIfNotExists(databasePtr, queryCreateTable)
}
