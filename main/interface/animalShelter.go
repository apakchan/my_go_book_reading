package main

import (
	"fmt"
	"math/rand"
)

type animalBehaviour interface {
	move() string
	eat() string
}

type animalType struct {
	animalName      string
	animalAvgHeight float64
	animalAvgWeight float64
	isDomestic      bool
	moveType        string
	food            string
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
	return fmt.Sprintf("%s %s %s", d.a.animalName, d.sb.move(), d.a.moveType)
}

func (d animal) eat() string {
	return fmt.Sprintf("%s %s %s", d.a.animalName, d.sb.eat(), d.a.food)
}

func main() {
	animals := []animal{
		{
			a: animalType{
				animalName: "dog",
				moveType:   "run",
				food:       "bone",
			},
		},
		{
			a: animalType{
				animalName: "cat",
				moveType:   "jump",
				food:       "fish",
			},
		},
		{
			a: animalType{
				animalName: "tiger",
				moveType:   "run",
				food:       "meat",
			},
		},
	}
	for i := 0; i < 3; i++ {
		for hour := 1; hour <= 24; hour++ {
			animal := animals[rand.Intn(len(animals))]
			if isWhiteDay(hour) {
				if rand.Intn(2) == 1 {
					fmt.Println(animal.eat())
				} else {
					fmt.Println(animal.move())
				}
			} else {
				fmt.Printf("%s sleep\n", animal.a.animalName)
			}
		}
	}
}

func isWhiteDay(hour int) bool {
	return hour >= 6 && hour <= 20
}
