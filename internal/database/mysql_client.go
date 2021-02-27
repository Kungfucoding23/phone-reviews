package database

import (
	"database/sql"

	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/logs"
	_ "github.com/go-sql-driver/mysql"
)

//MySQLClient ..
type MySQLClient struct {
	*sql.DB
}

//NewSQLClient sql client
func NewSQLClient(source string) *MySQLClient {
	db, err := sql.Open("mysql", source)
	if err != nil {
		logs.Log().Errorf("cannot create db tentat: %s", err.Error())
		panic(err)
	}
	//Ping a la db
	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to mysql")
	}
	return &MySQLClient{db}
}

//ViewStats ...DBStats contains database statistics.
func (c *MySQLClient) ViewStats() sql.DBStats {
	return c.Stats()
}
