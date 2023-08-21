package main

import "fmt"

var t interface {
	talk() string
}

type Martian struct{}

func (m Martian) talk() string {
	return "Martian"
}

type Laser int

func (l Laser) talk() string {
	return "Laser"
}

type starship struct {
	Laser
}

func main() {
	t = Martian{}
	fmt.Println(t.talk())
	t = Laser(1)
	fmt.Println(t.talk())
	s := starship{Laser(1)}
	s.talk()
}
