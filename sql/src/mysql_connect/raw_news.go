package mysqlconnect

import (
	"database/sql"
	"log"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func raw_news_insert_query(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("INSERT INTO raw_news(url,content,source) VALUES (?,?,?)")
	checkError(err)
	return stmt

}

func raw_news_insert_run(db *sql.DB, params ...interface{}) {

	stmt := raw_news_insert_query(db)

	defer stmt.Close()

	url := params[0].(string)
	news := params[1].(string)
	source := params[2].(string)

	result, err := stmt.Exec(url, news, source)
	checkError(err)
	rows, err := result.RowsAffected()
	checkError(err)
	log.Println(rows)

}

func Raw_news_insert(params ...interface{}) {

	Mysql_connect(raw_news_insert_run, params...)
}
