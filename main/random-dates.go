package main

import (
	"fmt"
	"math/rand"
)

const era = "AD"

type dateStruct struct {
	era   string
	year  int
	month int
	day   int
}

func isRun(year int) bool {
	return year%4 == 0 && year%100 != 0
}

func main() {
	for i := 0; i < 10; i++ {
		genDate()
	}
}

func genDate() {
	year := rand.Intn(2023-1949) + 1949 // [1949, 2023]
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		if isRun(year) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
