package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/atomic", newIdGenHandler(newIdGeneratorAtomic()))
	http.Handle("/mutex", newIdGenHandler(newIdGeneratorMutex()))
	http.Handle("/chan", newIdGenHandler(newIdGeneratorChan()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
