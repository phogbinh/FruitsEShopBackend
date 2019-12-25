package database_util

import (
	"database/sql"
	"errors"
	"strconv"

	. "backend/model"
	"backend/util"
)

// CreateTableIfNotExists creates a table using the given query if the table does not exist.
func CreateTableIfNotExists(databasePtr *sql.DB, query string) error {
	_, createTableError := databasePtr.Exec(query)
	return createTableError
}

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

// GetObjectByKeyColumn returns the object queried by the given key column.
func GetObjectByKeyColumn(queryGetObjectByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) ([]interface{}, Status) {
	objects, getStatus := GetObjectsByKeyColumn(queryGetObjectByKeyColumn, keyColumnValue, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return nil, getStatus
	}
	if len(objects) != 1 {
		return nil, util.StatusInternalServerError(GetObjectByKeyColumn, errors.New("Query 1 object but got "+strconv.Itoa(len(objects))+" object(s) instead."))
	}
	return objects[0], util.StatusOK()
}

// GetObjectsByKeyColumn returns the objects queried by the given key column.
func GetObjectsByKeyColumn(queryGetObjectsByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) ([][]interface{}, Status) {
	queryRowsPtr, queryError := databasePtr.Query(queryGetObjectsByKeyColumn, keyColumnValue)
	if queryError != nil {
		return nil, util.StatusInternalServerError(GetObjectsByKeyColumn, queryError)
	}
	defer queryRowsPtr.Close()
	return GetAllObjects(queryRowsPtr)
}

// GetAllObjects returns the objects fetched from the given database objects table rows pointer.
func GetAllObjects(databaseObjectsTableRowsPtr *sql.Rows) ([][]interface{}, Status) {
	columnNames, columnsError := databaseObjectsTableRowsPtr.Columns()
	if columnsError != nil {
		return nil, util.StatusInternalServerError(GetAllObjects, columnsError)
	}
	var objects [][]interface{}
	for databaseObjectsTableRowsPtr.Next() {
		object := make([]interface{}, len(columnNames))
		for index := range columnNames {
			object[index] = new(sql.RawBytes)
		}
		scanError := databaseObjectsTableRowsPtr.Scan(object...)
		if scanError != nil {
			return nil, util.StatusInternalServerError(GetAllObjects, scanError)
		}
		objects = append(objects, object)
	}
	return objects, util.StatusOK()
}

// GetRawBytesList returns a list of raw bytes from the given object.
func GetRawBytesList(object []interface{}) ([]sql.RawBytes, Status) {
	rawBytesList := make([]sql.RawBytes, len(object))
	for index := range rawBytesList {
		rawBytesPtr, ok := object[index].(*sql.RawBytes)
		if !ok {
			return nil, util.StatusInternalServerError(GetRawBytesList, errors.New("Error converting object entry to raw bytes pointer."))
		}
		rawBytesList[index] = *rawBytesPtr
	}
	return rawBytesList, util.StatusOK()
}
