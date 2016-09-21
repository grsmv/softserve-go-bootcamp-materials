package main

func main() {
	type a int
	type b int
	var z a
	var x b
	x = 1
	z = a(x)
	_ = z
	println("Hello world")
}
