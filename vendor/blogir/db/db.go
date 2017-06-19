package db

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var SQL *sql.DB

func createSchema() error {
    schema := `
        CREATE TABLE IF NOT EXISTS posts (
            title       VARCHAR(255),
            created_at  DATETIME,
            updated_at  DATETIME,
            body        TEXT
        );`

    _, err := SQL.Exec(schema)
    if err != nil {
        return err
    }

    return nil
}

func Connect(dbURL string) {
    var err error
    SQL, err = sql.Open("mysql", dbURL)
    if err != nil {
        log.Fatal("SQL driver error: ", err.Error())
    }

    if err = SQL.Ping(); err != nil {
        log.Fatal("Couldn't connect to DB: ", err.Error())
    }

    if err = createSchema(); err != nil {
        log.Fatal("Couldn't create DB schema: ", err.Error())
    }
}
