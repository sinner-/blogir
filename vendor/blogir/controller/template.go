package controller

import (
    "log"
    "fmt"
    "html/template"
    "net/http"
    "blogir/model"
)

var (
    funcMap = template.FuncMap{"toHTML": toHTML}
    templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
)

func toHTML(html string) template.HTML {
  return template.HTML(html)
}

func renderPage(w http.ResponseWriter, tmpl string, p *model.Page) {
    err := templates.ExecuteTemplate(w, fmt.Sprintf("%s.html", tmpl), p)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "HTTP 500", http.StatusInternalServerError)
    }
}

func renderIndex(w http.ResponseWriter, i *model.Index) {
    err := templates.ExecuteTemplate(w, "index.html", i)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "HTTP 500", http.StatusInternalServerError)
    }
}
