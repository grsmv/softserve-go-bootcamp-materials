package main

type Vocabulary map[string]string

func main() {
	var vocabulary = Vocabulary{
		"Ua": "Привіт, світ!",
		"En": "Hello world!",
		"De": "Hallo welt!",
	}

	println(vocabulary["Ua"])
}
