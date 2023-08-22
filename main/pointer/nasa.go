package main

import "fmt"

func main() {
	var administrator *string
	scolese := "scolese"
	bolden := "Charles F. Bolden"

	administrator = &scolese
	fmt.Printf("%s %T\n", *administrator, administrator)

	administrator = &bolden
	fmt.Printf("%s %T\n", *administrator, administrator)

	*administrator = "123"
	fmt.Println(bolden) // 123

	const a = "a"
	
}
