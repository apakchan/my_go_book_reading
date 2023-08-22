package main

import "fmt"

type animalBehaviour interface {
	move() string
	eat() string
}

type animalType struct {
	animalName      string
	animalAvgHeight float64
	animalAvgWeight float64
	isDomestic      bool
}

type simpleBehaviour struct {
}

func (sb simpleBehaviour) move() string {
	return "move"
}

func (sb simpleBehaviour) eat() string {
	return "eat"
}

type animal struct {
	a  animalType
	sb simpleBehaviour
}

func (d animal) move() string {
	return fmt.Sprintf("%s %s", d.a.animalName, d.sb.move())
}

func (d animal) eat() string {
	return fmt.Sprintf("%s %s", d.a.animalName, d.sb.eat())
}

func main() {
	dog := animal{
		a: animalType{
			animalName: "dog",
		},
	}
	fmt.Println(dog.move(), dog.eat())
}
