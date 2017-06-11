package main

import (
    "fmt"
    "blogir/db"
    "blogir/controller"
)

func main() {
    hostname := "127.0.0.1"
    port := 8080
    db.Connect()
    fmt.Println(fmt.Sprintf("Now listening on %s:%d", hostname, port))
    defer db.SQL.Close()
    controller.Start(hostname, port)
}
