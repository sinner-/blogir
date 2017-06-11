package controller

import (
    "net/http"
)

func loadRoutes() {
    http.Handle("/view/", pipeline(viewHandler))
    http.Handle("/edit/", pipeline(editHandler))
    http.Handle("/save/", pipeline(saveHandler))
}
