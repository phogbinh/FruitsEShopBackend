package database

import (
	. "backend/model"
	"backend/util"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func AddCartToUser(userName string, databasePtr *sql.DB) Status {
	_, insertError := databasePtr.Exec("INSERT INTO	" + CartTableName + " VALUES ();")
	if insertError != nil {
		return util.StatusInternalServerError(AddCartToUser, insertError)
	}
	_, insertError = databasePtr.Exec("INSERT INTO " + CustomerOwnCartTableName + " VALUES (" +
		"'" + userName + "', " + "(SELECT * FROM " + CartTableName + " ORDER BY " +
		CartIdColumnName + " DESC LIMIT 0 , 1));")
	if insertError != nil {
		return util.StatusInternalServerError(AddCartToUser, insertError)
	}
	return util.StatusOK()
}
