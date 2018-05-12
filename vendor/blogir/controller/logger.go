package controller

import (
    "net/http"
    "log"
)

func logger(r * http.Request) {
    log.Printf("%s %s %s - %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
}
