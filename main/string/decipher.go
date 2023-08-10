package main

import (
	"fmt"
)

// 根据维吉尼亚加密法
// 破解字符串
func main() {
	unKeyTheString()
}

func unKeyTheString() {
	plainText := "your message goes here"
	keyWord := "GOLANG"
	keyWordIndex := 0
	for _, character := range plainText {
		keyCharacter := keyWord[keyWordIndex]
		offset := keyCharacter - 'A'
		if !IsEnglish(character) {
			fmt.Printf("%c", character)
			continue
		}
		character = character - int32(offset)
		if character < 'a' {
			character += 26
		}
		fmt.Printf("%c", character)
		keyWordIndex = (keyWordIndex + 1) % len(keyWord)
	}
}

func keyTheString() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyWord := "GOLANG"
	keyWordIndex := 0
	for _, character := range cipherText {
		keyCharacter := keyWord[keyWordIndex]
		offset := keyCharacter - 'A'
		character = character - int32(offset)
		if character < 'A' {
			character += 26
		}
		fmt.Printf("%c", character)
		keyWordIndex = (keyWordIndex + 1) % len(keyWord)
	}
}
