package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

type authInfo struct{
    login string
    password string
}
func authMe(auth authInfo, h http.HandleFunc()) (http.HandleFunc) {
    return func(w http.ResponseWriter, r *http.Request){
	auth = strings.Split(r.Header.Get("Authorization"), " ")
	hasAuth := false
	authType := ""
	authToken := ""

	if len(auth) > 1 {
		authType = auth[0]
		authToken = auth[1]
		hasAuth = true
	}

	credentilas, err := base64.StdEncoding.DecodeString(authToken)
	authinfo := strings.Split(string(credentilas), ":")
	if authInfo[0] != auth.login || authInfo[1] != auth.pass {
		w.WriteHeader(403)
		w.Write([]byte("Authentification error"))
		return
	}
    h(w,r)
}
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Println("Start listening at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
