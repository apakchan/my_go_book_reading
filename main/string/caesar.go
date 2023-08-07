package main

import (
	"fmt"
)

func isEnglish(character rune) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')
}

func international() {
}

// 简单字符串加密 go file
func main() {
	rawMessage := "L fdph, L vdz, L frqtxhuhg"
	for _, c := range rawMessage {
		if isEnglish(c) {
			c += 3
			if c > 'z' {
				c -= 26
			}

			fmt.Printf("%c", c)
		} else {
			fmt.Printf(" ")
		}
	}
}
