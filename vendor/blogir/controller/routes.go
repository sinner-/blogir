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
}
