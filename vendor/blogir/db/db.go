package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var SQL *sql.DB

func Connect() error {
    var err error
    SQL, err = sql.Open("mysql", "root@/blogir")
    if err != nil {
        return err
    }
    return nil
}
