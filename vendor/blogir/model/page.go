package model

import (
    "blogir/db"
)

type Page struct {
    Title string
    Body string
    Recent []string
}

func (p *Page) Save() error {

    _, err := LoadPage(p.Title)
    if err == nil {
        _, err = db.SQL.Exec("UPDATE posts SET body = ? WHERE title = ?", p.Body , p.Title)
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
    row, err := db.SQL.Query("SELECT body FROM posts WHERE title = ?", title)
    if err != nil {
        return nil, err
    }
    defer row.Close()

    page := new(Page)
    page.Title = title

    row.Next()
    err = row.Scan(&page.Body)
    if err != nil {
        return nil, err
    }

    return page, nil
}
