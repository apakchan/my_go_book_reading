package main

import (
	"fmt"
	"math/rand"
)

func genNumber(floor int, celling int) int {
	if floor > celling {
		return -1
	}
	return rand.Intn(celling-floor) + floor
}

func main() {
	const speed = 100800 // km per hour
	const distance = 96300000
	fmt.Printf("SpaceX will spend %d hour going to Mars about %d day\n", distance/speed, distance/speed/24)
	fmt.Println(genNumber(56000000, 401000000))
	fmt.Println(rand.Intn(1))
}
