package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Create all database except users
func CreateDatabases(databasePtr *sql.DB) {
	createDatabaseAdminTableIfNotExists(databasePtr)
	createDatabaseProductTableIfNotExists(databasePtr)
	createDatabaseQATableIfNotExitsts(databasePtr)
	createDatabaseCartTableIfNotExists(databasePtr)
	createDatabaseActivityTableIfNotExists(databasePtr)
	createDatabaseAddToCartTableIfNotExists(databasePtr)
	createDatabaseBuyTableIfNotExists(databasePtr)
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
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS admin (\n" +
		"	username		VARCHAR(16) 	NOT NULL,\n" +
		"	password		VARCHAR(16)		NOT NULL,\n" +
		"	nickname		VARCHAR(15) 	NOT NULL DEFAULT 'user',\n" +
		"	email 			VARCHAR(36) 	NOT NULL,\n" +
		"	PRIMARY KEY (username),\n" +
		"	UNIQUE (email));")
	panicCreateTableError(createTableError)
}

func createDatabaseProductTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS product (\n" +
		"	p_id			INTEGER 		NOT NULL,\n" +
		"	s_username		VARCHAR(16) 	NOT NULL,\n" +
		"	description 	VARCHAR(255) 	DEFAULT 'No description',\n" +
		"	p_name			VARCHAR(255)	NOT NULL,\n" +
		"	category		VARCHAR(255)	DEFAULT 'None',\n" +
		"	source			VARCHAR(255)	DEFAULT 'None',\n" +
		"	price			INTEGER	 		NOT NULL,\n" +
		"	inventory		INTEGER			NOT NULL,\n" +
		"	sold_quantity	INTEGER			NOT NULL,\n" +
		"	onsale_date 	DATE			NOT NULL,\n" +
		"	PRIMARY KEY (p_id),\n" +
		"	FOREIGN KEY (s_username) REFERENCES users (username),\n" +
		"	CONSTRAINT p_id_non_negative			CHECK (p_id >= 0),\n" +
		"	CONSTRAINT price_non_negative 			CHECK (price >= 0),\n" +
		"	CONSTRAINT inventory_non_negative 		CHECK (inventory >= 0),\n" +
		"	CONSTRAINT sold_quantity_non_negative	CHECK (sold_quantity >= 0));")
	panicCreateTableError(createTableError)
}

func createDatabaseQATableIfNotExitsts(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS qa (\n" +
		"	p_id			INTEGER			NOT NULL,\n" +
		"	s_username		VARCHAR(16)		NOT NULL,\n" +
		"	c_username		VARCHAR(16)		NOT NULL,\n" +
		"	question		VARCHAR(100)	DEFAULT(''),\n" +
		"	answer			VARCHAR(100)	DEFAULT(''),\n" +
		"	ask_date		DATE			NOT NULL,\n" +
		"	ask_time		TIME			NOT NULL,\n" +
		"	ans_date		DATE,\n" +
		"	ans_time		TIME,\n" +
		"	PRIMARY KEY(p_id),\n" +
		"	FOREIGN KEY(p_id) REFERENCES product(p_id),\n" +
		"	FOREIGN KEY(s_username) REFERENCES users(username),\n" +
		"	FOREIGN KEY(c_username)	REFERENCES users(username),\n" +
		"	CONSTRAINT qa_check_date_interval		CHECK (ask_date < ans_date));")
	panicCreateTableError(createTableError)
}

func createDatabaseCartTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS cart (\n" +
		"	cart_id			INTEGER			NOT NULL,\n" +
		"	PRIMARY KEY(cart_id),\n" +
		"	CONSTRAINT cart_id_non_negative 		CHECK (cart_id >= 0));")
	panicCreateTableError(createTableError)
}

func createDatabaseActivityTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS activity (\n" +
		"	a_id			INTEGER			NOT NULL,\n" +
		"	name			VARCHAR(255)	NOT NULL,\n" +
		"	start_date		DATE			NOT NULL,\n" +
		"	end_date		DATE			NOT NULL,\n" +
		"	lowest_discount	FLOAT(2)		NOT NULL,\n" +
		"	PRIMARY KEY(a_id),\n" +
		"	CONSTRAINT activity_check_date_interval	CHECK (start_date < end_date),\n" +
		"	CONSTRAINT activity_discount_digit		CHECK (lowest_discount < 10));")
	panicCreateTableError(createTableError)
}

// ----------------------------Association table----------------------------
func createDatabaseAddToCartTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS add_to_cart (\n" +
		"	cart_id			INTEGER			NOT NULL,\n" +
		"	p_id			INTEGER			NOT NULL,\n" +
		"	quantity		INTEGER			NOT NULL,\n" +
		"	PRIMARY KEY(cart_id, p_id),\n" +
		"	FOREIGN KEY(cart_id) REFERENCES cart(cart_id),\n" +
		"	FOREIGN KEY(p_id) REFERENCES product(p_id));")
	panicCreateTableError(createTableError)
}

func createDatabaseBuyTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS buy (\n" +
		"	c_username		VARCHAR(16)		NOT NULL,\n" +
		"	cart_id			INTEGER			NOT NULL,\n" +
		"	buy_date		DATE			NOT NULL,\n" +
		"	PRIMARY KEY(c_username, cart_id),\n" +
		"	FOREIGN KEY(c_username) REFERENCES users(username),\n" +
		"	FOREIGN KEY(cart_id) REFERENCES cart(cart_id));")
	panicCreateTableError(createTableError)
}

func createDatabaseCEvaluateTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS c_evaluate (\n" +
		"	c_username		VARCHAR(16)		NOT NULL,\n" +
		"	s_username		VARCHAR(16)		NOT NULL,\n" +
		"	feedback		TEXT			NOT NULL,\n" +
		"	PRIMARY KEY(c_username, s_username),\n" +
		"	FOREIGN KEY(c_username) REFERENCES users(username),\n" +
		"	FOREIGN KEY(s_username) REFERENCES users(username));")
	panicCreateTableError(createTableError)
}

func createDatabaseSEvaluateTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS s_evaluate (\n" +
		"	s_username		VARCHAR(16)		NOT NULL,\n" +
		"	c_username		VARCHAR(16)		NOT NULL,\n" +
		"	feedback		TEXT			NOT NULL,\n" +
		"	PRIMARY KEY(s_username, c_username),\n" +
		"	FOREIGN KEY(s_username) REFERENCES users(username),\n" +
		"	FOREIGN KEY(c_username) REFERENCES users(username));")
	panicCreateTableError(createTableError)
}

func createDatabasePEvaluateTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS p_evaluate (\n" +
		"	p_id			INTEGER			NOT NULL,\n" +
		"	c_username		VARCHAR(16)		NOT NULL,\n" +
		"	feedback		TEXT			NOT NULL,\n" +
		"	PRIMARY KEY(p_id, c_username),\n" +
		"	FOREIGN KEY(p_id) REFERENCES product(p_id),\n" +
		"	FOREIGN KEY(c_username) REFERENCES users(username));")
	panicCreateTableError(createTableError)
}

func createDatabaseManageTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS manage (\n" +
		"	a_username		VARCHAR(16)		NOT NULL,\n" +
		"	username		VARCHAR(16)		NOT NULL,\n" +
		"	PRIMARY KEY(a_username, username),\n" +
		"	FOREIGN KEY(a_username) REFERENCES admin(username),\n" +
		"	FOREIGN KEY(username) REFERENCES users(username));")
	panicCreateTableError(createTableError)
}

func createDatabaseHoldTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS hold (\n" +
		"	a_username		VARCHAR(16)		NOT NULL,\n" +
		"	a_id			INTEGER			NOT NULL,\n" +
		"	PRIMARY KEY(a_username, a_id),\n" +
		"	FOREIGN KEY(a_username) REFERENCES admin(username),\n" +
		"	FOREIGN KEY(a_id) REFERENCES activity(a_id));")
	panicCreateTableError(createTableError)
}

func createDatabaseSJoinTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS s_join (\n" +
		"	s_username		VARCHAR(16)		NOT NULL,\n" +
		"	a_id			INTEGER			NOT NULL,\n" +
		"	PRIMARY KEY(s_username, a_id),\n" +
		"	FOREIGN KEY(s_username) REFERENCES users(username),\n" +
		"	FOREIGN KEY(a_id) REFERENCES activity(a_id));")
	panicCreateTableError(createTableError)
}

// Have the constraint that discount < activity.lowest_discount, but hard to describe in MySQL
func createDatabasePJoinTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS p_join (\n" +
		"	p_id			INTEGER			NOT NULL,\n" +
		"	a_id			INTEGER			NOT NULL,\n" +
		"	quantity		INTEGER			NOT NULL,\n" +
		"	discount		FLOAT(2)		NOT NULL,\n" +
		"	PRIMARY KEY(p_id, a_id),\n" +
		"	FOREIGN KEY(p_id) REFERENCES product(p_id),\n" +
		"	FOREIGN KEY(a_id) REFERENCES activity(a_id),\n" +
		"	CONSTRAINT pjoin_discount_digit			CHECK (discount < 10));")
	panicCreateTableError(createTableError)
}

func createDatabaseTakeOffTableIfNotExists(databasePtr *sql.DB) {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS take_off (\n" +
		"	a_username		VARCHAR(16)		NOT NULL,\n" +
		"	p_id			INTEGER			NOT NULL,\n" +
		"	PRIMARY KEY(a_username, p_id),\n" +
		"	FOREIGN KEY(a_username) REFERENCES admin(username),\n" +
		"	FOREIGN KEY(p_id) REFERENCES product(p_id));")
	panicCreateTableError(createTableError)
}

// panic error
func panicCreateTableError(err error) {
	if err != nil {
		panic(err)
	}
}
