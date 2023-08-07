package main

import "fmt"

// rune 表示单个统一码代码点, 是 int32 的别名
// rune 可以和 int32 互换, 用户可以自定义类型别名 be like:
// type byte = uint8
// type rune = int32
func main() {
	var pi rune = 960
	var alpha rune = 940
	var bang byte = 33
	fmt.Printf("%[1]v %[1]v %[1]v\n", pi, alpha, bang)
}
