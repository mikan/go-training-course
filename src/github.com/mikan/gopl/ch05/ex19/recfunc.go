// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("recFunc() returns " + strconv.Itoa(recFunc()))
}

func recFunc() (r int) {
	defer func() {
		p := recover()
		if i, ok := p.(int); ok { // using type assertion
			r = i
		} else {
			panic(p)
		}
	}()
	panic(1) // panic with non-zero value
}