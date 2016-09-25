package main

import (
	"time"
)

func do(num int) {
	println("Task ", num)
	time.Sleep(2 * time.Second)
}

func main() {
	for i := 1; i < 10; i++ {
		go do(i)
	}
}
