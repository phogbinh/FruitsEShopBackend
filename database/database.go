package database

import (
	"backend/database_users_table_util"
	"database/sql"

	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	productsTable "backend/database_products_table_util"

	_ "github.com/go-sql-driver/mysql"
)

// Create all database except users
func CreateDatabases(databasePtr *sql.DB) {
	createDatabaseAdminTableIfNotExists(databasePtr)
	createDatabaseProductTableIfNotExists(databasePtr)
	createDatabaseQATableIfNotExitsts(databasePtr)
	createDatabaseCartTableIfNotExists(databasePtr)
	createDatabaseActivityTableIfNotExists(databasePtr)
	createDatabaseOrderItemTableIfNotExists(databasePtr)
	createDatabaseCustomerOwnCartTableIfNotExists(databasePtr)
	createDatabaseTradeTableIfNotExists(databasePtr)
	createDatabaseCEvaluateTableIfNotExists(databasePtr)
	createDatabaseSEvaluateTableIfNotExists(databasePtr)
	createDatabasePEvaluateTableIfNotExists(databasePtr)
	createDatabaseManageTableIfNotExists(databasePtr)
	createDatabaseHoldTableIfNotExists(databasePtr)
	createDatabaseSJoinTableIfNotExists(databasePtr)
	createDatabasePJoinTableIfNotExists(databasePtr)
	createDatabaseTakeOffTableIfNotExists(databasePtr)
}

// ----------------------------Entity table----------------------------
func createDatabaseAdminTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + AdminTableName + " (\n" +
		AdminUserNameColumnName + " VARCHAR(30) 	NOT NULL,\n" +
		AdminPasswordColumnName + " VARCHAR(30)		NOT NULL,\n" +
		AdminNicknameColumnName + "	VARCHAR(30) 	NOT NULL DEFAULT 'user',\n" +
		AdminEmailColumnName + "	VARCHAR(320) 	NOT NULL,\n" +
		"PRIMARY KEY (" + AdminUserNameColumnName + "),\n" +
		"UNIQUE (" + AdminEmailColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseProductTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + ProductTableName + " (\n" +
		ProductIdColumnName + " 			INTEGER 		NOT NULL,\n" +
		ProductStaffUserNameColumnName + " 	VARCHAR(30) 	NOT NULL,\n" +
		ProductDescriptionColumnName + "	VARCHAR(255) 	DEFAULT 'No description',\n" +
		ProductNameColumnName + "			VARCHAR(255)	NOT NULL,\n" +
		ProductCategoryColumnName + "		VARCHAR(255)	DEFAULT 'None',\n" +
		ProductSourceColumnName + "			VARCHAR(255)	DEFAULT 'None',\n" +
		ProductPriceColumnName + "			INTEGER	 		NOT NULL,\n" +
		ProductInventoryColumnName + "		INTEGER			NOT NULL,\n" +
		ProductSoldQuantityColumnName + "	INTEGER			NOT NULL,\n" +
		ProductOnSaleDataColumnName + "		DATE			NOT NULL,\n" +
		productsTable.SpecialEventDiscountPolicyCodeColumnName + "	CHAR(9),\n" +
		"PRIMARY KEY (" + ProductIdColumnName + "),\n" +
		"FOREIGN KEY (" + ProductStaffUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + " (" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + productsTable.SpecialEventDiscountPolicyCodeColumnName + ") REFERENCES " + discountPoliciesTablesConst.SpecialEventDiscountPoliciesTableName + "(" + discountPoliciesTablesConst.SpecialEventDiscountPoliciesCodeColumnName + ")\n" +
		"	ON DELETE SET NULL,\n" +
		"CONSTRAINT p_id_non_negative			CHECK (" + ProductIdColumnName + " >= 0),\n" +
		"CONSTRAINT price_non_negative 			CHECK (" + ProductPriceColumnName + " >= 0),\n" +
		"CONSTRAINT inventory_non_negative 		CHECK (" + ProductInventoryColumnName + " >= 0),\n" +
		"CONSTRAINT sold_quantity_non_negative	CHECK (" + ProductSoldQuantityColumnName + " >= 0));")
	panicCreateTableError(createTableError)
}

func createDatabaseQATableIfNotExitsts(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + QA_TableName + " (\n" +
		QA_ProductIdColumnName + "			INTEGER			NOT NULL,\n" +
		QA_StaffUserNameColumnName + "		VARCHAR(30)		NOT NULL,\n" +
		QA_CustomerUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		QA_QuestionColumnname + "			VARCHAR(100)	DEFAULT(''),\n" +
		QA_AnswerColumnName + "				VARCHAR(100)	DEFAULT(''),\n" +
		QA_AskDatetimeColumnName + "		DATETIME		NOT NULL,\n" +
		QA_AnsDatetimeColumnName + "		DATETIME,\n" +
		"PRIMARY KEY(" + QA_ProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + QA_ProductIdColumnName + ") REFERENCES " + ProductTableName + "(" + ProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + QA_StaffUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + QA_CustomerUserNameColumnName + ")	REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"CONSTRAINT qa_check_date_interval		CHECK (" + QA_AskDatetimeColumnName + " < " + QA_AnsDatetimeColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseCartTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + CartTableName + " (\n" +
		CartIdColumnName + "	INTEGER			NOT NULL	AUTO_INCREMENT,\n" +
		"PRIMARY KEY(" + CartIdColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseActivityTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + ActivityTableName + " (\n" +
		ActivityIdColumnName + "				INTEGER			NOT NULL,\n" +
		ActivityNameColumnName + "				VARCHAR(255)	NOT NULL,\n" +
		ActivityStartDateColumnName + "			DATE			NOT NULL,\n" +
		ActivityEndDateColumnName + "			DATE			NOT NULL,\n" +
		ActivityLowestDiscountColumnName + "	FLOAT(2)		NOT NULL,\n" +
		"PRIMARY KEY(" + ActivityIdColumnName + "),\n" +
		"CONSTRAINT activity_check_date_interval	CHECK (" + ActivityStartDateColumnName + " < " + ActivityEndDateColumnName + "),\n" +
		"CONSTRAINT activity_discount_digit		CHECK (" + ActivityLowestDiscountColumnName + " < 10));")
	panicCreateTableError(createTableError)
}

// ----------------------------Association table----------------------------
func createDatabaseOrderItemTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + OrderItemTableName + " (\n" +
		OrderItemCartIdColumnName + "		INTEGER			NOT NULL,\n" +
		OrderItemProductIdColumnName + "	INTEGER			NOT NULL,\n" +
		OrderItemQuantity + "				INTEGER			NOT NULL,\n" +
		"PRIMARY KEY(" + OrderItemCartIdColumnName + ", " + OrderItemProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + OrderItemCartIdColumnName + ") REFERENCES " + CartTableName + "(" + CartIdColumnName + "),\n" +
		"FOREIGN KEY(" + OrderItemProductIdColumnName + ") REFERENCES " + ProductTableName + "(" + ProductIdColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseCustomerOwnCartTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + CustomerOwnCartTableName + " (\n" +
		CustomerOwnCartCustomerUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		CustomerOwnCartCartIdColumnName + "				INTEGER		NOT NULL,\n" +
		"PRIMARY KEY(" + CustomerOwnCartCustomerUserNameColumnName + ", " + CustomerOwnCartCartIdColumnName + "),\n" +
		"FOREIGN KEY(" + CustomerOwnCartCustomerUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + CustomerOwnCartCartIdColumnName + ") REFERENCES " + CartTableName + "(" + CartIdColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseTradeTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + TradeTableName + " (\n" +
		TradeCartIdColumnName + "			INTEGER		NOT NULL,\n" +
		TradeProductIdColumnName + "		INTEGER		NOT NULL,\n" +
		TradeProductQuantityColumnName + "	INTEGER		NOT NULL,\n" +
		TradeDateTimeColumnName + "			DATETIME	NOT NULL,\n" +
		"PRIMARY KEY(" + TradeCartIdColumnName + ", " + TradeProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + TradeProductIdColumnName + ") REFERENCES " + ProductTableName + "(" + ProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + TradeCartIdColumnName + ") REFERENCES " + CartTableName + "(" + CartIdColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseCEvaluateTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + CustomerEvaluateTableName + " (\n" +
		CustomerEvaluateCustomerUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		CustomerEvaluateStaffUserNameColumnName + "		VARCHAR(30)		NOT NULL,\n" +
		CustomerEvaluateFeedbackColumnName + "			TEXT			NOT NULL,\n" +
		"PRIMARY KEY(" + CustomerEvaluateCustomerUserNameColumnName + ", " + CustomerEvaluateStaffUserNameColumnName + "),\n" +
		"FOREIGN KEY(" + CustomerEvaluateCustomerUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + CustomerEvaluateStaffUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseSEvaluateTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + StaffEvaluateTableName + " (\n" +
		StaffEvaluateCustomerUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		StaffEvaluateStaffUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		StaffEvaluateFeedbackColumnName + "			TEXT			NOT NULL,\n" +
		"PRIMARY KEY(" + StaffEvaluateCustomerUserNameColumnName + ", " + StaffEvaluateStaffUserNameColumnName + "),\n" +
		"FOREIGN KEY(" + StaffEvaluateCustomerUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + StaffEvaluateStaffUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabasePEvaluateTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + ProductEvaluateTableName + " (\n" +
		ProductEvaluateProductIdColumnName + "			INTEGER			NOT NULL,\n" +
		ProductEvaluateCustomerUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		ProductEvaluateFeedbackColumnName + "			TEXT			NOT NULL,\n" +
		"PRIMARY KEY(" + ProductEvaluateProductIdColumnName + ", " + ProductEvaluateCustomerUserNameColumnName + "),\n" +
		"FOREIGN KEY(" + ProductEvaluateProductIdColumnName + ") REFERENCES " + ProductTableName + "(" + ProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + ProductEvaluateCustomerUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseManageTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + ManageTableName + " (\n" +
		ManageAdminUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		ManageUserNameColumnName + "		VARCHAR(30)		NOT NULL,\n" +
		"PRIMARY KEY(" + ManageAdminUserNameColumnName + ", " + ManageUserNameColumnName + "),\n" +
		"FOREIGN KEY(" + ManageAdminUserNameColumnName + ") REFERENCES " + AdminTableName + "(" + AdminUserNameColumnName + "),\n" +
		"FOREIGN KEY(" + ManageUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseHoldTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + HoldTableName + " (\n" +
		HoldAdminUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		HoldActivityIdColumnName + "	INTEGER			NOT NULL,\n" +
		"PRIMARY KEY(" + HoldAdminUserNameColumnName + ", " + HoldActivityIdColumnName + "),\n" +
		"FOREIGN KEY(" + HoldAdminUserNameColumnName + ") REFERENCES " + AdminTableName + "(" + AdminUserNameColumnName + "),\n" +
		"FOREIGN KEY(" + HoldActivityIdColumnName + ") REFERENCES " + ActivityTableName + "(" + ActivityIdColumnName + "));")
	panicCreateTableError(createTableError)
}

func createDatabaseSJoinTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + StaffJoinActivityTableName + " (\n" +
		StaffJoinActivityStaffUserNameColumnName + "	VARCHAR(30)		NOT NULL,\n" +
		StaffJoinActivityActivityIdColumnName + "		INTEGER			NOT NULL,\n" +
		"PRIMARY KEY(" + StaffJoinActivityStaffUserNameColumnName + ", " + StaffJoinActivityActivityIdColumnName + "),\n" +
		"FOREIGN KEY(" + StaffJoinActivityStaffUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + StaffJoinActivityActivityIdColumnName + ") REFERENCES " + ActivityTableName + "(" + ActivityIdColumnName + "));")
	panicCreateTableError(createTableError)
}

// Have the constraint that discount < activity.lowest_discount, but hard to describe in MySQL
func createDatabasePJoinTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + ProductJoinActivityTableName + " (\n" +
		ProductJoinActivityProductIdColumnName + "	INTEGER			NOT NULL,\n" +
		ProductJoinActivityActivityIdColumnName + "	INTEGER			NOT NULL,\n" +
		ProductJoinActivityQuantityColumnName + "	INTEGER			NOT NULL,\n" +
		ProductJoinActivityDiscountColumnName + "	FLOAT(2)		NOT NULL,\n" +
		"PRIMARY KEY(" + ProductJoinActivityProductIdColumnName + ", " + ProductJoinActivityActivityIdColumnName + "),\n" +
		"FOREIGN KEY(" + ProductJoinActivityProductIdColumnName + ") REFERENCES " + ProductTableName + "(" + ProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + ProductJoinActivityActivityIdColumnName + ") REFERENCES " + ActivityTableName + "(" + ActivityIdColumnName + "),\n" +
		"CONSTRAINT pjoin_discount_digit			CHECK (" + ProductJoinActivityDiscountColumnName + " < 10));")
	panicCreateTableError(createTableError)
}

func createDatabaseTakeOffTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + TakeOffTableName + " (\n" +
		TakeOffUserNameColumnName + "		VARCHAR(30)		NOT NULL,\n" +
		TakeOffProductIdColumnName + "		INTEGER			NOT NULL,\n" +
		"PRIMARY KEY(" + TakeOffUserNameColumnName + ", " + TakeOffProductIdColumnName + "),\n" +
		"FOREIGN KEY(" + TakeOffUserNameColumnName + ") REFERENCES " + database_users_table_util.TableName + "(" + database_users_table_util.UserNameColumnName + "),\n" +
		"FOREIGN KEY(" + TakeOffProductIdColumnName + ") REFERENCES " + ProductTableName + "(" + ProductIdColumnName + "));")
	panicCreateTableError(createTableError)
}

// panic error
func panicCreateTableError(err error) {
	if err != nil {
		panic(err)
	}
}
