// Package database provides a main database connection
package database

import "database/sql"

// InitDB initialize database connection to mysql.
func InitDB() *sql.DB {
	connectionString := "root@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql",connectionString)
	if err != nil {
		panic(err.Error())
	}
	return databaseConnection
}
