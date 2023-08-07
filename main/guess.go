package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const guessRes = 50
	i := 1
	var a = rand.Intn(100) + 1
	for a <= guessRes {
		a = rand.Intn(100) + 1
		i++
	}
	fmt.Println(a)
	fmt.Println(i)
}
