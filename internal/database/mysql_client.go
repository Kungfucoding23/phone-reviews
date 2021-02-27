package database

import (
	"database/sql"
	"fmt"
)

//MySQLClient ..
type MySQLClient struct {
}

//NewSQLClient sql client
func NewSQLClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)
	if err != nil {
		_ = fmt.Errorf("cannot create db tentat: %s", err.Error())
		panic(err)
	}
	return db
}
