package model

import (
    "time"
    "blogir/db"
)

type Entry struct {
    Title         string
    Created       string
}

type Index struct {
  Entries []*Entry
  Authenticated bool
  Name string
}

var (
  timeLayout = "2006-01-02 15:04:05"
)

func LoadIndex() (*Index, error) {
    rows, err := db.SQL.Query("SELECT title, created_at FROM posts ORDER BY created_at DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    index := new(Index)

    for rows.Next() {
        entry := new(Entry)
        err = rows.Scan(&entry.Title, &entry.Created)
        t, _ := time.Parse(timeLayout, entry.Created)
        entry.Created = t.Format("Monday, 02 Jan 2006")

        if err != nil {
            return nil, err
        }

        index.Entries = append(index.Entries, entry)
    }

    return index, nil
}
