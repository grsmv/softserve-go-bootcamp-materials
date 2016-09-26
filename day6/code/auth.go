package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type authInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func AuthHandlerFunc(auth authInfo, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequest authInfo
		log.Println("Trying to authenticate")
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&authRequest)
		if err != nil {
			log.Println("Can't parse payload: ", err)
			w.WriteHeader(500)
			w.Write([]byte("Can't parse payload"))
			return
		}
		if authRequest != auth {
			log.Printf("Wrong login and password. Login: %q, Password: %q", authRequest.Login, authRequest.Password)
			w.WriteHeader(403)
			w.Write([]byte("Authentification error"))
			return
		}
		h(w, r)
		log.Println("Successfully executed handler")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Restricted area, but you are authorized!")
}

func main() {
	auth := authInfo{Login: "admin", Password: "password"}
	http.HandleFunc("/", AuthHandlerFunc(auth, handler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
