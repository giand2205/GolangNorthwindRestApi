package database

import (
	"database/sql"
	"fmt"
)

func InitDB() *sql.DB{
	connectionString := "root:sasa@tcp(localhost:3306)/northwind"
	fmt.Println(connectionString)
	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}
	return databaseConnection
}
