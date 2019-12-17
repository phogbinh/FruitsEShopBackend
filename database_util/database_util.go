package database_util

import (
	"database/sql"

	. "backend/model"
	"backend/util"
)

// PrepareThenExecuteQuery prepares the query and executes it using the given execute arguments.
func PrepareThenExecuteQuery(databasePtr *sql.DB, query string, executeArguments ...interface{}) Status {
	prepareStatementPtr, prepareError := databasePtr.Prepare(query)
	if prepareError != nil {
		return util.StatusInternalServerError(PrepareThenExecuteQuery, prepareError)
	}
	_, executeError := prepareStatementPtr.Exec(executeArguments...)
	if executeError != nil {
		return util.StatusInternalServerError(PrepareThenExecuteQuery, executeError)
	}
	return util.StatusOK()
}
