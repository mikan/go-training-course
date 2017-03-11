// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"time"

	"github.com/mikan/go-training-course/ch02/ex03/popcount"
	"github.com/mikan/go-training-course/ch02/ex03/popcount2"
)

func main() {
	ch := make(chan string)
	go loop(1, 10000000, 1023, ch)
	go loop(2, 10000000, 1023, ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func loop(method int, count int, input uint64, ch chan<- string) {
	start := time.Now()
	for i := 1; i < count; i++ {
		switch method {
		case 1:
			popcount.PopCount(input)
		case 2:
			popcount2.PopCount(input)
		default:
			panic("unknown method")
		}
	}
	secs := time.Since(start).Seconds()
	if ch != nil {
		ch <- fmt.Sprintf("[popcount%d] Elapsed time: %.2fs (count = %d)", method, secs, count)
	}
}

// Result (popcount2 is fucking slow.):
// [popcount1] Elapsed time: 0.06s (count = 10000000)
// [popcount2] Elapsed time: 0.19s (count = 10000000)
