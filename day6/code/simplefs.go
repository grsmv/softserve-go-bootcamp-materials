package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Println("Start listening at :3000")
	http.ListenAndServe(":3000", nil)
}
