package utils

import "database/sql"

func GetDBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/cinema-demo.db")
	if err != nil {
		panic(err)
	}

	return db
}
