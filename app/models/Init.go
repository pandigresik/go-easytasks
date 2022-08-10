package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbCon *sql.DB
	err error
)

func InitDB(filepath string) {
	dbCon , err = sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if dbCon == nil {
		panic(nil)
	}
	migrate()
}

func migrate() {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		pic VARCHAR NOT NULL,
		deadline VARCHAR NOT NULL,
		status INTEGER
    );
    `

	_, err := dbCon.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}