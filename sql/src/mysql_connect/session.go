package mysqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql_connect() *sql.DB {

	dbUser := os.Getenv("dbUser") // DB username
	dbPass := os.Getenv("dbPass") // DB Password
	dbHost := os.Getenv("dbHost") // DB Hostname/IP
	dbName := os.Getenv("dbName") // Database name

	db_url := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", db_url)
	checkError(err)

	log.Println("Successfully connected to the db")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db

}
