package main

import "fmt"

type convertCF func(float64) float64

func drawTable(cf convertCF) {
	fmt.Println("================================================================")
	if cf(0) > 0 {
		fmt.Println("| °C                            | °F                           |")
	} else {
		fmt.Println("| °F                            | °C                           |")
	}
	fmt.Println("================================================================")
	for i := -40; i <= 100; i += 5 {
		tar := cf(float64(i))
		fmt.Printf("| %3d                           | %6.2f                       |\n", i, tar)
	}
	fmt.Println("================================================================")
}

func main() {
	drawTable(func(f float64) float64 {
		return f*9/5 + 32
	})
	drawTable(func(f float64) float64 {
		return (f - 32) * 5 / 9
	})
}
