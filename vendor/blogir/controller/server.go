package controller

import (
    "fmt"
    "log"
    "net/http"
    "blogir/db"
)

func Start() {
    log.Print("Loading configuration...")
    loadConfig()

    log.Print("Connecting to database...")
    db.Connect(CONF.dbURL)
    defer db.SQL.Close()

    log.Print("Loading HTTP routes...")
    loadRoutes()

    log.Print("Server listening on ", CONF.listenURL)
    http.ListenAndServe(fmt.Sprintf("%s", CONF.listenURL), nil)
}
