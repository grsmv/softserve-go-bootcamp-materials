package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("This is a HTTPS example server.\n"))
	})
	log.Fatal(http.ListenAndServeTLS(":4443", "certificate.pem", "private.key", nil))
}
