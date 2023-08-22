package main

import "fmt"

type turtle struct {
	x int `json:"x"`
	y int `json:"y"`
}

func (t *turtle) up(offset int) {
	move(t, offset, "up")
}

func (t *turtle) down(offset int) {
	move(t, offset, "down")
}

func (t *turtle) left(offset int) {
	move(t, offset, "left")
}

func (t *turtle) right(offset int) {
	move(t, offset, "right")
}

func move(t *turtle, offset int, methodType string) {
	if methodType == "up" {
		t.y = t.y + offset
	}
	if methodType == "down" {
		t.y = t.y - offset
	}
	if methodType == "left" {
		t.x = t.x - offset
	}
	if methodType == "right" {
		t.x = t.x + offset
	}
}

func main() {
	t := turtle{
		x: 2,
		y: 1,
	}
	t.left(2)
	t.down(1)
	fmt.Printf("%d %d\n", t.x, t.y)
}
