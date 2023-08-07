package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = math.Pi
	third := 1.0 / 3
	testNum := 0.1
	testNum += 0.2
	fmt.Println(testNum)
	fmt.Println(third)
	fmt.Println(PI)
	fmt.Printf("%0015.3f\n", third)
	fmt.Println(third * 3)
}
