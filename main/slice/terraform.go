package main

import "fmt"

type Planets string

func (p Planets) terraform() Planets {
	if p == "Mars" || p == "Uranus" || p == "Neptune" {
		p = "New" + p
	}
	return p
}

func main() {
	planets := []Planets{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	for i, ele := range planets {
		planets[i] = ele.terraform()
	}
	fmt.Println(planets)
}
