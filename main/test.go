package main

import "fmt"

func main() {
	var stringVar = "abc"
	switch stringVar {
	case "abc":
		fmt.Println("abc")
		fallthrough
	case "bdc":
		fmt.Println("bdc")
		fallthrough // 执行下一个分支代码
	default:
		fmt.Println("def")
	}
}
