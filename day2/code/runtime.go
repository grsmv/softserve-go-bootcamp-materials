package main

import "fmt"

func process(j *int) {
	fmt.Printf("Execption: %+v\n\n", recover())
	fmt.Printf("Index: %d\n", *j)
}

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var i int
	i = -1
	defer process(&i)
	for i = 0; i <= len(arr); i++ {
		println(arr[i])
	}

	println(arr[i])
}
