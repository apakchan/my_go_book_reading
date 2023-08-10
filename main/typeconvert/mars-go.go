package main

import "fmt"

// calculate my age in mars
func main() {
	age := 22
	earthDays := 365.2425
	marsDays := 687.0
	marAge := float64(age)
	marAge = marAge * earthDays / marsDays
	fmt.Printf("mars age: %.0f", marAge)

}
