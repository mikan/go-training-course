// Copyright 2016 mikan. All rights reserved.

package main

import "fmt"

const (
	KB = float64(1000)
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("1KB = %.0f\n", KB)
	fmt.Printf("1MB = %.0f\n", MB)
	fmt.Printf("1GB = %.0f\n", GB)
	fmt.Printf("1TB = %.0f\n", TB)
	fmt.Printf("1PB = %.0f\n", PB)
	fmt.Printf("1EB = %.0f\n", EB)
	fmt.Printf("1ZB = %.0f\n", ZB)
	fmt.Printf("1YB = %.0f\n", YB)
}

// Output:
// 1KB = 1000
// 1MB = 1000000
// 1GB = 1000000000
// 1TB = 1000000000000
// 1PB = 1000000000000000
// 1EB = 1000000000000000000
// 1ZB = 1000000000000000000000
// 1YB = 999999999999999983222784
