package main

import (
	"fmt"
	"strconv"
)

func main() {
	countDown := 10
	fmt.Println("Launch in T minus " + strconv.Itoa(countDown) + " seconds.")
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Printf("parse err!\n")
	}
	println(parseBool)
	atoi, err := strconv.Atoi("s0")
	if err != nil {
		fmt.Println("atoi err!")
	}
	fmt.Println(atoi)
	//intToString := fmt.Sprintln(123)
	println(strconv.ParseBool("0"))
}
