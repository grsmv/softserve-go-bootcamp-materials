package main

import (
	j "encoding/json"
	"fmt"
)

type Vocabulary map[string]string

func main() {
	var vocabulary = Vocabulary{
		"Ua": "Привіт, світ!",
		"En": "Hello world!",
		"De": "Hallo welt!",
	}
	v, _ := j.Marshal(vocabulary)
	fmt.Print(string(v))
	println(vocabulary["Ua"])
}
