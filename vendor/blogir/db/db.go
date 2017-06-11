package db

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var SQL *sql.DB

func Connect(dbURL string) {
    var err error
    SQL, err = sql.Open("mysql", dbURL)
    if err != nil {
        log.Fatal("SQL driver error: ", err.Error())
    }

    if err = SQL.Ping(); err != nil {
        log.Fatal("Couldn't connect to DB: ", err.Error())
    }
}
