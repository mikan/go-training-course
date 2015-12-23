// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program name is " + os.Args[0])
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}