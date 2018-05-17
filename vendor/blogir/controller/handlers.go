package controller

import (
    "fmt"
    "errors"
    "net/http"
    "log"
    "regexp"
    "crypto/hmac"
    "crypto/sha512"
    "encoding/base64"
    "golang.org/x/crypto/bcrypt"
    "blogir/model"
)

var validTitle = regexp.MustCompile("^[A-Za-z0-9 -]+$")
var validFile = regexp.MustCompile("((css|img|html)/[A-Za-z]+\\.(css|jpg|html))$")

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

func getFile(url string) (string, error) {
    file := url[len("/static/"):]
    if !validFile.MatchString(file) {
        return "", errors.New("Requested invalid static file.")
    }
    return file, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    i, err := model.LoadIndex()
    if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
    }

    i.Authenticated = isAuthenticated(r)

    renderIndex(w, i)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
    file, err := getFile(r.URL.Path)
    if err != nil {
        log.Print(err.Error())
        return
    }
    http.ServeFile(w, r, fmt.Sprintf("static/%s", file))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(r.URL.Path, "/view/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    p, err := model.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, fmt.Sprintf("/admin/edit/%s", title), http.StatusFound)
        return
    }

    p.Authenticated = isAuthenticated(r)

    renderPage(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(r.URL.Path, "/admin/save/")
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
    title, err := getTitle(r.URL.Path, "/admin/edit/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    p, err := model.LoadPage(title)
    if err != nil {
        p = &model.Page{Title: title}
    }
    renderPage(w, "edit", p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  if r.FormValue("username") == CONF.adminUsername && bcrypt.CompareHashAndPassword(CONF.adminPasshash, []byte(r.FormValue("password"))) == nil {
      hash := hmac.New(sha512.New, CONF.cookieSigningKey)
      hash.Write(CONF.adminCookieString)
      cookie := http.Cookie{Name: "auth", Value: base64.StdEncoding.EncodeToString(hash.Sum(nil))}
      http.SetCookie(w, &cookie)
      http.Redirect(w, r, "/", http.StatusFound)
  } else {
      http.Error(w, "Bad login.", http.StatusForbidden)
  }
}

func newHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, fmt.Sprintf("/admin/edit/%s", r.FormValue("title")), http.StatusFound)
}
