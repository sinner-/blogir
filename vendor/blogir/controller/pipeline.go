package controller

import (
    "net/http"
)

func pipeline(h http.HandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logger(r)
        serveAuthenticated(w, r)
        h.ServeHTTP(w, r)
    })
}
