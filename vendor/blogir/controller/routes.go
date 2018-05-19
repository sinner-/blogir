package controller

import (
    "net/http"
)

func loadRoutes() {
    http.Handle("/", pipeline(indexHandler))
    http.Handle("/login", pipeline(loginHandler))
    http.Handle("/static/", pipeline(staticHandler))
    http.Handle("/view/", pipeline(viewHandler))
    http.Handle("/admin/edit/", pipeline(editHandler))
    http.Handle("/admin/save/", pipeline(saveHandler))
    http.Handle("/admin/new", pipeline(newHandler))
    http.Handle("/admin/delete/", pipeline(deleteHandler))
}
