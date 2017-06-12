package controller

import (
    "fmt"
    "errors"
    "net/http"
    "log"
    "regexp"
    "blogir/model"
)

var validTitle = regexp.MustCompile("([A-Za-z0-9]+)$")

func getTitle(url string, path string) (string, error) {
    title := url[len(path):]
    if len(title) < 1 {
        return "", errors.New("No title specified.")
    }
    if !validTitle.MatchString(title) {
        return "", errors.New("Malformed title.")
    }
    return title, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(r.URL.Path, "/edit/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    p, err := model.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, fmt.Sprintf("/edit/%s", title), http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(r.URL.Path, "/edit/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    body := r.FormValue("body")
    p := &model.Page{Title: title, Body: body}
    err = p.Save()
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "HTTP 500", http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, fmt.Sprintf("/view/%s", title), http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(r.URL.Path, "/edit/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    p, err := model.LoadPage(title)
    if err != nil {
        p = &model.Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

