package main

import (
	"fmt"
	"math/rand"
	"time"
)

const width = 80
const height = 15

type Universe [][]bool

func NewUniverse() Universe {
	res := make(Universe, 0, 15)
	for i := 0; i < 15; i++ {
		res = append(res, make([]bool, 80))
	}
	return res
}

func (u Universe) Show() {
	for _, row := range u {
		for _, status := range row {
			if status {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

/*
Seed 激活大约 25% 的细胞,  20% - 30% 的范围
*/
func generateNumSeq(rangeNum int, totalNum int) []int {
	finalNum := totalNum * rangeNum / 100
	resIntSlice := make([]int, 0, finalNum)
	filter := map[int]bool{}
	for finalNum > 0 {
		genNum := rand.Intn(totalNum)
		if filter[genNum] {
			continue
		}
		filter[genNum] = true
		resIntSlice = append(resIntSlice, genNum)
		finalNum--
	}
	return resIntSlice
}

func (u Universe) Seed() Universe {
	rangeNum := 30 - rand.Intn(11)
	totalRow, totalColumn := cap(u), cap(u[0])
	seq := generateNumSeq(rangeNum, totalRow*totalColumn)
	for _, num := range seq {
		row, column := num/totalColumn, num%totalColumn
		u[row][column] = true
	}
	return u
}

/*
NextAlive
1. 当一个存活细胞附近的存活细胞少于 2 个时, 死亡
2. 当一个存活细胞附近有 2-3 个存活细胞, 延续下一代
3. 当一个存活细胞附近有多于 3 个存活细胞, 死亡
4. 当一个死亡细胞附近正好有 3 个存活细胞, 该死亡细胞复活
*/
func (u Universe) NextAlive(x, y int) bool {
	totalRow, totalColumn := cap(u), cap(u[0])
	aroundLive := 0
	for cursorX := -1; cursorX <= 1; cursorX++ {
		for cursorY := -1; cursorY <= 1; cursorY++ {
			if cursorX == 0 && cursorY == 0 {
				continue
			}
			jumpX, jumpY := x+cursorX, y+cursorY
			if jumpX < 0 {
				jumpX += totalRow
			}
			if jumpY < 0 {
				jumpY += totalColumn
			}
			jumpX, jumpY = jumpX%totalRow, jumpY%totalColumn
			if u[jumpX][jumpY] {
				aroundLive++
			}
		}
	}
	return (u[x][y] && aroundLive <= 3 && aroundLive >= 2) || (aroundLive == 3)
}

func Step(a, b Universe) {
	for x, eachRow := range a {
		for y, _ := range eachRow {
			isAlive := a.NextAlive(x, y)
			if isAlive {
				b[x][y] = true
			}
		}
	}
	for x, eachRow := range b {
		for y, _ := range eachRow {
			a[x][y] = b[x][y]
		}
	}
}

func (u Universe) countAlive() int {
	allAlive := 0
	for _, eachRow := range u {
		for _, ele := range eachRow {
			if ele {
				allAlive++
			}
		}
	}
	return allAlive
}

func main() {
	a := NewUniverse().Seed()
	for i := 0; i < 100; i++ {
		a.Show()
		fmt.Printf("allLive: %d\n", a.countAlive())
		for i := 0; i < 80; i++ {
			fmt.Printf("-")
		}
		fmt.Println()

		Step(a, NewUniverse())
		time.Sleep(2 * time.Second)
	}
}
