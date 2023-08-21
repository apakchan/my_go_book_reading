package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	shouter := strings.ToUpper(t.talk())
	fmt.Println(shouter)
}

func main() {

}
