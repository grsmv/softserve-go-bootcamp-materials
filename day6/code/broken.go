package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Random panic
	if rand.Intn(3) == 1 {
		panic("Boom!")
	}
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
