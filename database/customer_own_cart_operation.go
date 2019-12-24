package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Add order product to cart
func GetCartIdWithUsername(userName string, databasePtr *sql.DB) (code int, cartId int) {
	row, err := databasePtr.Query("SELECT "+CartIdColumnName+"\n"+
		"	FROM	"+CustomerOwnCartTableName+"\n"+
		"	WHERE	"+CustomerOwnCartCustomerUserNameColumnName+" = ?;", userName)

	if err != nil {
		code = 400
		return code, -1
	} else {
		code = 200
	}

	row.Next()
	err = row.Scan(&cartId)
	fmt.Println("username = ", userName)
	fmt.Println("cart id = ", cartId)
	if err != nil {
		code = 403
		return code, -1
	}

	return code, cartId
}
