package model

import (
    "blogir/db"
)

type Entry struct {
    Title         string
    Created       string
    Updated       string
}

type Index struct {
  Entries []*Entry
  Authenticated bool
}

func LoadIndex() (*Index, error) {
    rows, err := db.SQL.Query("SELECT title, created_at, updated_at FROM posts ORDER BY created_at DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    index := new(Index)

    for rows.Next() {
      entry := new(Entry)
      err = rows.Scan(&entry.Title, &entry.Created, &entry.Updated)

      err := rows.Scan(&entry.Title, &entry.Created, &entry.Updated)
      if err != nil {
          return nil, err
      }

      index.Entries = append(index.Entries, entry)
    }

    return index, nil
}
