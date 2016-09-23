package main

import (
	"sync"
	"time"
)

func do(num int, wg *sync.WaitGroup) {
	println("Task ", num)
	time.Sleep(2 * time.Second)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go do(i, &wg)
	}
	wg.Wait()
}
