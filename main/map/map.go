package main

import "fmt"

func main() {
	// key: string, value: int
	temp := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temp["Earth"], temp["abc"])
	temp["Earth"] = 30
	fmt.Println(temp["Earth"])
	if non, ok := temp["non"]; ok {
		fmt.Println(non)
	} else {
		fmt.Println("non doesnt exist")
	}
}
