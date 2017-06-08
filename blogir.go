package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
    "log"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
    log.Print(r.URL.Path)
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(fmt.Sprintf("%s.html", tmpl))
    t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    log.Print(r.URL.Path)
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, fmt.Sprintf("/edit/%s", title), http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    log.Print(r.URL.Path)
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: body}
    p.save()
    http.Redirect(w, r, fmt.Sprintf("/view/%s", title), http.StatusFound)
}

type Page struct {
    Title string
    Body string
}

func (p *Page) save() error {
    db, err := sql.Open("mysql", "root@/blogir")
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM posts WHERE title = ?", p.Title)
    if err != nil {
        return err
    }

    _, err = db.Exec("INSERT INTO posts VALUES (?, ?)", p.Title, p.Body)
    if err != nil {
        return err
    }

    return nil
}

func loadPage(title string) (*Page, error) {
    db, err := sql.Open("mysql", "root@/blogir")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    row, err := db.Query("SELECT body FROM posts WHERE title = ?", title)
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

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    hostname := "127.0.0.1"
    port := 8080
    connstring := fmt.Sprintf("%s:%d", hostname, port)
    fmt.Println("Now listening on", connstring)
    http.ListenAndServe(connstring, nil)
}
