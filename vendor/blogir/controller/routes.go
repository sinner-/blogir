package controller

import (
    "net/http"
)

func loadRoutes() {
    http.Handle("/", pipeline(indexHandler))
    http.Handle("/static/", pipeline(staticHandler))
    http.Handle("/view/", pipeline(viewHandler))
    http.Handle("/edit/", pipeline(editHandler))
    http.Handle("/save/", pipeline(saveHandler))
}
