package controller

import (
    "html/template"
    "net/http"
    "log"
    "blogir/model"
    "fmt"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
    t, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", tmpl))
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "HTTP 500", http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, p)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "HTTP 500", http.StatusInternalServerError)
    }
}


