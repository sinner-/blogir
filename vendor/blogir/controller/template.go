package controller

import (
    "log"
    "fmt"
    "html/template"
    "net/http"
    "blogir/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderPage(w http.ResponseWriter, tmpl string, p *model.Page) {
    err := templates.ExecuteTemplate(w, fmt.Sprintf("%s.html", tmpl), p)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "HTTP 500", http.StatusInternalServerError)
    }
}
