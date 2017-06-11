package controller

import (
    "net/http"
    "log"
    "fmt"
)

func logger(r * http.Request) {
    log.Print(fmt.Sprintf("%s %s - %s", r.Proto, r.Method, r.URL))
}

