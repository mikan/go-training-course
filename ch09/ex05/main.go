package main

import (
	"fmt"
	"time"
)

func pingPong(ping, pong chan int) {
	for msg := range ping {
		pong <- msg + 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)
	go pingPong(ping, pong)
	go pingPong(pong, ping)
	ping <- 1
	<-time.NewTimer(1 * time.Second).C
	fmt.Println(<-pong/2, "ping/sec")
	// Result on MacBook Early 2016:
	// 1454679 ping/sec
	// 1399950 ping/sec
	// 1460891 ping/sec
}
