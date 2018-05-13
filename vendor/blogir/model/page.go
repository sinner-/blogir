package model

import (
    "blogir/db"
)

type Page struct {
    Title         string
    Body          string
    Recent        []string
    Created       string
    Updated       string
    Authenticated bool
}

func (p *Page) Save() error {

    _, err := LoadPage(p.Title)
    if err == nil {
        _, err = db.SQL.Exec("UPDATE posts SET body = ?, updated_at = NOW() WHERE title = ?", p.Body , p.Title)
        if err != nil {
            return err
        }
    } else {
        _, err = db.SQL.Exec("INSERT INTO posts VALUES (?, NOW(), NOW(), ?)", p.Title, p.Body)
        if err != nil {
            return err
        }
    }

    return nil
}

func LoadPage(title string) (*Page, error) {
    row, err := db.SQL.Query("SELECT created_at, updated_at, body FROM posts WHERE title = ?", title)
    if err != nil {
        return nil, err
    }
    defer row.Close()

    page := new(Page)
    page.Title = title

    row.Next()
    err = row.Scan(&page.Created, &page.Updated, &page.Body)
    if err != nil {
        return nil, err
    }

    return page, nil
}
