package model

import (
    "blogir/db"
)

func LoadRecent() ([]string, error) {
    rows, err := db.SQL.Query("SELECT title FROM posts ORDER BY created_at DESC LIMIT 20")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    recent := []string{}
    row := ""

    for rows.Next() {
        err = rows.Scan(&row)
        if err != nil {
            return nil, err
        }
        recent = append(recent, row)
    }

    return recent, nil
}
