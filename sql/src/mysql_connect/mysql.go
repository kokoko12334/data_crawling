package mysqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type sql_query func(*sql.DB, ...interface{})

func Mysql_connect(logic sql_query, params ...interface{}) {

	var condition bool

	server := os.Getenv("YOUR_SERVER")
	fmt.Println(server)
	if server == "LOCAL" {
		condition = true
	} else {
		condition = false
	}

	if condition {
		DEV_DB_connection(logic, params...)
		// OP_DB_connection(logic)
	} else {
		// SSH_DEV_DB_connection(logic)
		// SSH_OP_DB_connection(logic)
		fmt.Println("NOT YET")
	}

}
func DEV_DB_connection(logic sql_query, params ...interface{}) {

	dbUser := os.Getenv("dbUser") // DB username
	dbPass := os.Getenv("dbPass") // DB Password
	dbHost := os.Getenv("dbHost") // DB Hostname/IP
	dbName := os.Getenv("dbName") // Database name

	db_url := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	if db, err := sql.Open("mysql", db_url); err == nil {
		fmt.Printf("Successfully connected to the db\n")

		logic(db, params...)

		defer db.Close()
	} else {
		fmt.Println(err)
	}

}
