package main

import (
	"time"
	"fmt"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball
}
