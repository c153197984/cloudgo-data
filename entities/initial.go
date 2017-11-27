package entities

import "database/sql"

import _ "github.com/go-sql-driver/mysql" // Load the driver of mysql.

const (
	dbPath string = "root:root@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=true"
)

var db *sql.DB

// An SQLExecutor encapsulates all functions that execute sql
// statements as an interface.
type SQLExecutor interface {
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*sql.Stmt, error)
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	QueryRow(sql string, args ...interface{}) *sql.Row
}

// An DataAccessObject is a data access object containing an
// SQLExecutor interface.
type DataAccessObject struct {
	SQLExecutor
}

func init() {
	var err error
	db, err = sql.Open("mysql", dbPath)
	errHandler(err)
}

func errHandler(err error) error {
	if err != nil {
		panic(err)
	}
	return err
}
