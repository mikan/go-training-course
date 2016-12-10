// Copyright 2015-2016 mikan. All rights reserved.

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

// Depends on tempconv package in the ex01
import (
	"fmt"
	"os"
	"strconv"

	"github.com/mikan/go-training-course/ch02/ex02/lengthconv"
	"github.com/mikan/go-training-course/ch02/ex02/weightconv"
	"github.com/mikan/go-training-course/ch02/ex01/tempconv"
)

func main() {
	var input = os.Args[1:]
	if len(os.Args[1:]) == 0 {
		fmt.Print("value > ")
		var value string
		fmt.Scan(&value)
		input = make([]string, 1)
		input[0] = value
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		showAsTemp(t)
		showAsLength(t)
		showAsWeight(t)
	}
}

func showAsTemp(temp float64) {
	f := tempconv.Fahrenheit(temp)
	c := tempconv.Celsius(temp)
	fmt.Printf("[TEMP] %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func showAsLength(length float64) {
	f := lengthconv.Feat(length)
	m := lengthconv.Meters(length)
	fmt.Printf("[LENGTH] %s = %s, %s = %s\n", f, lengthconv.FtoM(f), m, lengthconv.MtoF(m))
}

func showAsWeight(weight float64) {
	p := weightconv.Pounds(weight)
	kg := weightconv.Kilograms(weight)
	fmt.Printf("[WEIGHT] %s = %s, %s = %s\n", p, weightconv.PtoKG(p), kg, weightconv.KGtoP(kg))
}

// bin\ex02 10
// [TEMP] 10℉ = -12.222222222222221℃, 10℃ = 50℉
// [LENGTH] 10ft = 3.048m, 10m = 32.808ft
// [WEIGHT] 10lb = 4.535923700000001kg, 10kg = 22.046lb
