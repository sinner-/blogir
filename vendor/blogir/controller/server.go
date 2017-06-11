package controller

import (
    "fmt"
    "net/http"
)

func Start(hostname string, port int) {
    loadRoutes()
    http.ListenAndServe(fmt.Sprintf("%s:%d", hostname, port), nil)
}
