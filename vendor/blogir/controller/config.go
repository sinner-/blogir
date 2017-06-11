package controller

import (
    "fmt"
    "os"
)

type serverConfig struct {
    listenURL   string
    dbURL       string
}

var (
    CONF    serverConfig
)

func genListenURL() string {
    listenAddress, present := os.LookupEnv("BLOGIR_LISTEN_ADDRESS")
    if !present || listenAddress == "" {
        listenAddress = "127.0.0.1"
    }

    listenPort, present := os.LookupEnv("BLOGIR_LISTEN_PORT")
    if !present || listenPort == "" {
        listenPort = "8080"
    }

    return fmt.Sprintf("%s:%s", listenAddress, listenPort)
}

func genDBURL() string {
    dbUser, present := os.LookupEnv("BLOGIR_DB_USER")
    if !present || dbUser == "" {
        dbUser = "root"
    }

    dbPassword, present := os.LookupEnv("BLOGIR_DB_PASSWORD")
    if !present {
        dbPassword = ""
    }

    dbHost, present := os.LookupEnv("BLOGIR_DB_HOST")
    if !present || dbHost == "" {
        dbHost = "127.0.0.1"
    }

    dbPort, present := os.LookupEnv("BLOGIR_DB_PORT")
    if !present || dbPort == "" {
        dbPort = "3306"
    }

    dbName, present := os.LookupEnv("BLOGIR_DB_NAME")
    if !present || dbPort == "" {
        dbName = "blogir"
    }

    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func loadConfig() {
    CONF.listenURL = genListenURL()
    CONF.dbURL = genDBURL()
}
