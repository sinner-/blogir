package controller

import (
    "fmt"
    "os"
)

type serverConfig struct {
    blogName          string
    listenURL         string
    dbURL             string
    adminUsername     string
    adminPasshash     []byte
    cookieSigningKey  []byte
    adminCookieString []byte
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
    if !present {
        dbName = "blogir"
    }

    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func loadConfig() {
    CONF.listenURL = genListenURL()
    CONF.dbURL = genDBURL()

    adminUsername, present := os.LookupEnv("BLOGIR_ADMIN_USERNAME")
    if !present {
      adminUsername = "admin"
    }

    CONF.adminUsername = adminUsername

    adminPasshash, present := os.LookupEnv("BLOGIR_ADMIN_PASSHASH")
    if !present {
        fmt.Println("You must specify BLOGIR_ADMIN_PASSHASH environment variable.")
        os.Exit(1)
    }

    CONF.adminPasshash = []byte(adminPasshash)

    cookieSigningKey, present := os.LookupEnv("BLOGIR_COOKIE_SIGNING_KEY")
    if !present {
        fmt.Println("You must specify BLOGIR_COOKIE_SIGNING_KEY environment variable.")
        os.Exit(1)
    }

    CONF.cookieSigningKey = []byte(cookieSigningKey)

    adminCookieString, present := os.LookupEnv("BLOGIR_ADMIN_COOKIE_STRING")
    if !present {
        adminCookieString = "admin"
    }

    CONF.adminCookieString = []byte(adminCookieString)

    blogName, present := os.LookupEnv("BLOGIR_BLOG_NAME")
    if !present {
        blogName = "Sina's Blog"
    }

    CONF.blogName = blogName

}
