package main

type Vocabulary struct {
	Ua string
	En string
	De string
}

func main() {
	var vocabulary = Vocabulary{
		Ua: "Привіт, світ!",
		En: "Hello world!",
		De: "Hallo welt!",
	}

	println(vocabulary.De)
}
