package main

import (
	j "encoding/json"
	"fmt"
)

type Vocabulary struct {
	Ua string `json:"ua"`
	En string
	De string
}

func main() {
	var vocabulary = Vocabulary{
		Ua: "Привіт, світ!",
		En: "Hello world!",
		De: "Hallo welt!",
	}
	v, _ := j.Marshal(vocabulary)
	fmt.Println(string(v))
	println(vocabulary.De)
}
