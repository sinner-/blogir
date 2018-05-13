package controller

import (
    "net/http"
    "strings"
    "crypto/hmac"
    "crypto/sha512"
    "encoding/base64"
)

func isAuthenticated(r * http.Request) (bool) {
    cookie, err := r.Cookie("auth")
    if err != nil {
        return false
    }

    messageMAC, err := base64.StdEncoding.DecodeString(cookie.Value)
    if err != nil {
        return false
    }

    mac := hmac.New(sha512.New, CONF.cookieSigningKey)
    mac.Write(CONF.adminCookieString)
    expectedMAC := mac.Sum(nil)

    return hmac.Equal(messageMAC, expectedMAC)
}

func serveAuthenticated(w http.ResponseWriter, r * http.Request) {
    if strings.HasPrefix(r.URL.Path, "/admin") && !isAuthenticated(r) {
        http.Redirect(w, r, "/static/html/login.html", http.StatusFound)
    }
}
