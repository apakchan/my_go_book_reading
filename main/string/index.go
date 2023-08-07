package main

import "fmt"

func main() {
	message := "shalom"
	c := message[5] // m
	//message[5] = 'd'  无法给 go string 赋值, string 自诞生则不可变
	fmt.Println(c)
	fmt.Println(message)
}
