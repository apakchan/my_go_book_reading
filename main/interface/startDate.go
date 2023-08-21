package main

import (
	"fmt"
	"time"
)

type startDate interface {
	YearDay() int
	Hour() int
}

func starDate(date startDate) float64 {
	day := float64(date.YearDay())
	h := float64(date.Hour())
	return 1000 + day + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func main() {
	fmt.Println(starDate(sol(1234)))
	fmt.Println(starDate(time.Date(2023, 8, 22, 1, 1, 1, 0, time.UTC)))
}
