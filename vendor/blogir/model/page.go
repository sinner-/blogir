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
    _, err := db.SQL.Exec("DELETE FROM posts WHERE title = ?", p.Title)
    if err != nil {
        return err
    }

    _, err = db.SQL.Exec("INSERT INTO posts VALUES (?, NOW(), NOW(), ?)", p.Title, p.Body)
    if err != nil {
        return err
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

    page.Recent, err = loadRecent()
    if err != nil {
        return nil, err
    }

    row.Next()
    err = row.Scan(&page.Body)
    if err != nil {
        return nil, err
    }

    return page, nil
}
