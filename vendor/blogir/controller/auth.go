package controller

import (
    "net/http"
    "strings"
    "crypto/hmac"
    "crypto/sha512"
    "encoding/base64"
)

func isAuthenticated(w http.ResponseWriter, r * http.Request) (bool) {
    if !strings.HasPrefix(r.URL.Path, "/admin") {
        return true
    }

    cookie, err := r.Cookie("auth")
    if err != nil {
        http.Redirect(w, r, "/static/login.html", http.StatusFound)
        return false
    }

    messageMAC, err := base64.StdEncoding.DecodeString(cookie.Value)
    if err != nil {
        http.Redirect(w, r, "/static/login.html", http.StatusFound)
        return false
    }

    mac := hmac.New(sha512.New, CONF.cookieSigningKey)
    mac.Write(CONF.adminCookieString)
    expectedMAC := mac.Sum(nil)

    return hmac.Equal(messageMAC, expectedMAC)
}
