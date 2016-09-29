package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

type authInfo struct {
	login    string
	password string
}

func authMe(credentials authInfo, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.Split(r.Header.Get("Authorization"), " ")
		hasAuth := false
		authType := ""
		authToken := ""

		if len(auth) > 1 {
			authType = auth[0]
			authToken = auth[1]
			hasAuth = true
		}
		if !hasAuth || authType != "Basic" {
			w.Header().Add("WWW-Authenticate", `Basic realm="Protected area"`)
			w.WriteHeader(401)
			w.Write([]byte("Authentification required"))
			return
		}
		cred, err := base64.StdEncoding.DecodeString(authToken)
		if err != nil {
			w.WriteHeader(401)
			w.Header().Add("WWW-Authenticate", `Basic realm="Protected area"`)
			w.Write([]byte("Base authentification error"))
			return
		}

		authinfo := strings.Split(string(cred), ":")
		if len(authinfo) < 2 {
			w.WriteHeader(401)
			w.Header().Add("WWW-Authenticate", `Basic realm="Protected area"`)
			w.Write([]byte("Authentification error on split"))
			return
		}
		log.Println(authinfo[0], authinfo[1])
		if strings.Trim(authinfo[0], " \t\n") != credentials.login || strings.Trim(authinfo[1], " \t\n") != credentials.password {
			w.WriteHeader(403)
			w.Header().Add("WWW-Authenticate", `Basic realm="Protected area"`)
			w.Write([]byte("Authentification error"))
			return
		}
		h.ServeHTTP(w, r)
	}
}

func main() {
	valid := authInfo{"user", "pass"}
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", authMe(valid, fs))
	log.Println("Start listening at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
