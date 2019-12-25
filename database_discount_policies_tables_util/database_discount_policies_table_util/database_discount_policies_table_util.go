package database_discount_policies_table_util

import (
	"database/sql"
	"errors"
	"strconv"

	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	discountPolicyTypesTable "backend/database_discount_policies_tables_util/database_discount_policy_types_table_util"
	seasoningsDiscountPoliciesTable "backend/database_discount_policies_tables_util/database_seasonings_discount_policies_table_util"
	shippingDiscountPoliciesTable "backend/database_discount_policies_tables_util/database_shipping_discount_policies_table_util"
	specialEventDiscountPoliciesTable "backend/database_discount_policies_tables_util/database_special_event_discount_policies_table_util"
	usersTable "backend/database_users_table_util"
	"backend/database_util"
	. "backend/model"
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
	ShippingMinimumOrderPriceColumnName = "ShippingMinimumOrderPrice"
	SeasoningsRateColumnName            = "SeasoningsRate"
	SeasoningsBeginDateColumnName       = "SeasoningsBeginDate"
	SeasoningsEndDateColumnName         = "SeasoningsEndDate"
	SpecialEventRateColumnName          = "SpecialEventRate"
	SpecialEventBeginDateColumnName     = "SpecialEventBeginDate"
	SpecialEventEndDateColumnName       = "SpecialEventEndDate"
)

const (
	queryCreateTable = "CREATE TABLE IF NOT EXISTS " + tableName + util.EndOfLine +
		"(" + util.EndOfLine +
		codeColumnName + "			CHAR(9)			NOT NULL," + util.EndOfLine +
		nameColumnName + "			VARCHAR(255)	NOT NULL," + util.EndOfLine +
		descriptionColumnName + "	VARCHAR(255)	NOT NULL," + util.EndOfLine +
		typeColumnName + "			VARCHAR(255)	NOT NULL," + util.EndOfLine +
		staffUserNameColumnName + "	VARCHAR(30)		NOT NULL," + util.EndOfLine +
		"PRIMARY KEY(" + codeColumnName + ")," + util.EndOfLine +
		"FOREIGN KEY(" + typeColumnName + ") REFERENCES " + discountPoliciesTablesConst.DiscountPolicyTypesTableName + "(" + discountPoliciesTablesConst.DiscountPolicyTypesNameColumnName + ")" + util.EndOfLine +
		"	ON UPDATE CASCADE," + util.EndOfLine +
		"FOREIGN KEY(" + staffUserNameColumnName + ") REFERENCES " + usersTable.TableName + "(" + usersTable.UserNameColumnName + ")" + util.EndOfLine +
		"	ON DELETE CASCADE" + util.EndOfLine +
		")"
	queryInsertDiscountPolicy = "INSERT INTO " + tableName + " VALUES(?, ?, ?, ?, ?)"
	queryGetDiscountPolicies  = "SELECT" + util.EndOfLine +
		tableName + "." + codeColumnName + "," + util.EndOfLine +
		tableName + "." + nameColumnName + "," + util.EndOfLine +
		tableName + "." + descriptionColumnName + "," + util.EndOfLine +
		tableName + "." + typeColumnName + "," + util.EndOfLine +
		tableName + "." + staffUserNameColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.ShippingDiscountPoliciesTableName + "." + discountPoliciesTablesConst.ShippingDiscountPoliciesMinimumOrderPriceColumnName + "	AS " + ShippingMinimumOrderPriceColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.SeasoningsDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SeasoningsDiscountPoliciesRateColumnName + "			AS " + SeasoningsRateColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.SeasoningsDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SeasoningsDiscountPoliciesBeginDateColumnName + "		AS " + SeasoningsBeginDateColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.SeasoningsDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SeasoningsDiscountPoliciesEndDateColumnName + "			AS " + SeasoningsEndDateColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.SpecialEventDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SpecialEventDiscountPoliciesRateColumnName + "		AS " + SpecialEventRateColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.SpecialEventDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SpecialEventDiscountPoliciesBeginDateColumnName + "	AS " + SpecialEventBeginDateColumnName + "," + util.EndOfLine +
		discountPoliciesTablesConst.SpecialEventDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SpecialEventDiscountPoliciesEndDateColumnName + "		AS " + SpecialEventEndDateColumnName + util.EndOfLine +
		"FROM " + tableName + util.EndOfLine +
		"		LEFT OUTER JOIN " + discountPoliciesTablesConst.ShippingDiscountPoliciesTableName + "		ON " + tableName + "." + codeColumnName + " = " + discountPoliciesTablesConst.ShippingDiscountPoliciesTableName + "." + discountPoliciesTablesConst.ShippingDiscountPoliciesCodeColumnName + util.EndOfLine +
		"		LEFT OUTER JOIN " + discountPoliciesTablesConst.SeasoningsDiscountPoliciesTableName + "		ON " + tableName + "." + codeColumnName + " = " + discountPoliciesTablesConst.SeasoningsDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SeasoningsDiscountPoliciesCodeColumnName + util.EndOfLine +
		"		LEFT OUTER JOIN " + discountPoliciesTablesConst.SpecialEventDiscountPoliciesTableName + "	ON " + tableName + "." + codeColumnName + " = " + discountPoliciesTablesConst.SpecialEventDiscountPoliciesTableName + "." + discountPoliciesTablesConst.SpecialEventDiscountPoliciesCodeColumnName + util.EndOfLine
	queryGetDiscountPolicyByCode  = queryGetDiscountPolicies + "WHERE " + tableName + "." + codeColumnName + " = ?"
	queryGetStaffDiscountPolicies = queryGetDiscountPolicies + "WHERE " + tableName + "." + staffUserNameColumnName + " = ?"
	queryDeleteDiscountPolicy     = "DELETE FROM " + tableName + " WHERE " + codeColumnName + " = ?"
)

// CreateTableIfNotExists creates table `discount_policies`.
func CreateTableIfNotExists(databasePtr *sql.DB) error {
	return database_util.CreateTableIfNotExists(databasePtr, queryCreateTable)
}

// InsertDiscountPolicyToSuperclassAndSubclassTables inserts the given discount policy to the database superclass table `discount_policies` and to the subclass table of corresponding type.
func InsertDiscountPolicyToSuperclassAndSubclassTables(discountPolicy DiscountPolicy, databasePtr *sql.DB) Status {
	insertToSuperclassTableStatus := insertDiscountPolicy(discountPolicy, databasePtr)
	if !util.IsStatusOK(insertToSuperclassTableStatus) {
		return insertToSuperclassTableStatus
	}
	insertToSubclassTableStatus := insertDiscountPolicyToSubclassTable(discountPolicy, databasePtr)
	if !util.IsStatusOK(insertToSubclassTableStatus) {
		return insertToSubclassTableStatus
	}
	return util.StatusOK()
}

func insertDiscountPolicy(discountPolicy DiscountPolicy, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryInsertDiscountPolicy, discountPolicy.Code, discountPolicy.Name, discountPolicy.Description, discountPolicy.Type, discountPolicy.StaffUserName)
}

func insertDiscountPolicyToSubclassTable(discountPolicy DiscountPolicy, databasePtr *sql.DB) Status {
	if discountPolicy.Type == discountPolicyTypesTable.TypeShipping {
		return shippingDiscountPoliciesTable.InsertDiscountPolicy(discountPolicy, databasePtr)
	} else if discountPolicy.Type == discountPolicyTypesTable.TypeSeasonings {
		return seasoningsDiscountPoliciesTable.InsertDiscountPolicy(discountPolicy, databasePtr)
	} else if discountPolicy.Type == discountPolicyTypesTable.TypeSpecialEvent {
		return specialEventDiscountPoliciesTable.InsertDiscountPolicy(discountPolicy, databasePtr)
	} else {
		return util.StatusInternalServerError(insertDiscountPolicyToSubclassTable, errors.New("Invalid discount policy type."))
	}
}

// DeleteDiscountPolicy deletes the given discount policy from the database table `discount_policies` and cascadingly from related database tables.
func DeleteDiscountPolicy(discountPolicyCode string, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryDeleteDiscountPolicy, discountPolicyCode)
}

// IsExistingDiscountPolicy determines whether the given discount policy is in the database table `discount_policies`.
func IsExistingDiscountPolicy(discountPolicyCode string, databasePtr *sql.DB) (bool, Status) {
	discountPolicies, getStatus := getDiscountPoliciesByKeyColumn(queryGetDiscountPolicyByCode, discountPolicyCode, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return false, getStatus
	}
	return len(discountPolicies) > 0, util.StatusOK()
}

// GetStaffDiscountPolicies returns all discount policies' information by the given staff's user name.
func GetStaffDiscountPolicies(staffUserName string, databasePtr *sql.DB) ([]DiscountPolicy, Status) {
	return getDiscountPoliciesByKeyColumn(queryGetStaffDiscountPolicies, staffUserName, databasePtr)
}

// GetDiscountPolicyByCode returns a discount policy's information by the given discount policy code.
func GetDiscountPolicyByCode(code string, databasePtr *sql.DB) (DiscountPolicy, Status) {
	return getDiscountPolicyByKeyColumn(queryGetDiscountPolicyByCode, code, databasePtr)
}

func getDiscountPolicyByKeyColumn(queryGetDiscountPolicyByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) (DiscountPolicy, Status) {
	var dumpDiscountPolicy DiscountPolicy
	discountPolicies, getStatus := getDiscountPoliciesByKeyColumn(queryGetDiscountPolicyByKeyColumn, keyColumnValue, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return dumpDiscountPolicy, getStatus
	}
	if len(discountPolicies) != 1 {
		return dumpDiscountPolicy, util.StatusInternalServerError(getDiscountPolicyByKeyColumn, errors.New("Query 1 discount policy but got "+strconv.Itoa(len(discountPolicies))+" discount policy(s) instead."))
	}
	return discountPolicies[0], util.StatusOK()
}

func getDiscountPoliciesByKeyColumn(queryGetDiscountPoliciesByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) ([]DiscountPolicy, Status) {
	queryRowsPtr, queryError := databasePtr.Query(queryGetDiscountPoliciesByKeyColumn, keyColumnValue)
	if queryError != nil {
		return nil, util.StatusInternalServerError(getDiscountPoliciesByKeyColumn, queryError)
	}
	defer queryRowsPtr.Close()
	discountPolicies, getStatus := getAllDiscountPolicies(queryRowsPtr)
	if !util.IsStatusOK(getStatus) {
		return nil, getStatus
	}
	return discountPolicies, util.StatusOK()
}

func getAllDiscountPolicies(databaseDiscountPoliciesTableRowsPtr *sql.Rows) ([]DiscountPolicy, Status) {
	var discountPolicies []DiscountPolicy
	for databaseDiscountPoliciesTableRowsPtr.Next() {
		var discountPolicy DiscountPolicy
		scanError := databaseDiscountPoliciesTableRowsPtr.Scan(&discountPolicy.Code, &discountPolicy.Name, &discountPolicy.Description, &discountPolicy.Type, &discountPolicy.StaffUserName, &discountPolicy.ShippingMinimumOrderPrice, &discountPolicy.SeasoningsRate, &discountPolicy.SeasoningsBeginDate, &discountPolicy.SeasoningsEndDate, &discountPolicy.SpecialEventRate, &discountPolicy.SpecialEventBeginDate, &discountPolicy.SpecialEventEndDate)
		if scanError != nil {
			return nil, util.StatusInternalServerError(getAllDiscountPolicies, scanError)
		}
		discountPolicies = append(discountPolicies, discountPolicy)
	}
	return discountPolicies, util.StatusOK()
}
