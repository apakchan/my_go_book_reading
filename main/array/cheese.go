package main

import "fmt"

func printCheese(cheese [4][4]rune) {
	for _, arr := range cheese {
		for i, ele := range arr {
			if i < (len(arr) - 1) {
				fmt.Printf("'%c' ", ele)
			} else {
				fmt.Printf("'%c'", ele)
			}
		}
		fmt.Println()
	}
}

func main() {
	test := [4][4]rune{
		{'1', '2', '3', '4'},
		{'5', '6', '7', '8'},
		{'a', 'b', 'c', 'd'},
		{'e', 'f', 'g', 'h'},
	}
	printCheese(test)
}
