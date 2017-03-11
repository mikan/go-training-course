// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"time"

	"github.com/mikan/go-training-course/ch02/ex03/popcount"
	"github.com/mikan/go-training-course/ch02/ex03/popcount2"
	"github.com/mikan/go-training-course/ch02/ex04/popcount3"
	"github.com/mikan/go-training-course/ch02/ex05/popcount4"
)

func main() {
	ch := make(chan string)
	go loop(1, 10000000, 1023, ch)
	go loop(2, 10000000, 1023, ch)
	go loop(3, 10000000, 1023, ch)
	go loop(4, 10000000, 1023, ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
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
		case 3:
			popcount3.PopCount(input)
		case 4:
			popcount4.PopCount(input)
		default:
			panic("unknown method")
		}
	}
	secs := time.Since(start).Seconds()
	if ch != nil {
		ch <- fmt.Sprintf("[popcount%d] Elapsed time: %.2fs (count = %d)", method, secs, count)
	}
}

// Result:
// [popcount1] Elapsed time: 0.07s (count = 10000000)
// [popcount4] Elapsed time: 0.20s (count = 10000000)
// [popcount2] Elapsed time: 0.21s (count = 10000000)
// [popcount3] Elapsed time: 1.59s (count = 10000000)
