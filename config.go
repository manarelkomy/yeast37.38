package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GtDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "university"
	dbUser := "root"
	dbPass := "123456"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	return
}
