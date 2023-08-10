package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	// [0:4] 代表 [0, 4)
	terrestrial := planets[0:4]
	// [4:6] 代表 [4, 6)
	gasGiants := planets[4:6]
	// [6:] 代表 [6, len(arr))
	iceGiants := planets[6:]
	// [:] 代表数组全部元素
	allPlanets := planets[:]
	fmt.Println(terrestrial, gasGiants, iceGiants, allPlanets)
	// 如果一个切片是从一个数组中切出来的, 那么可以将这个切片看成是该数组的一个视图
	// 对切片中任意一个元素赋予新值都会导致原数组的改变
	terrestrial[0] = "new"
	println(terrestrial[0]) // new
	println(planets[0])     // new
	neptune := "Neptune"
	tune := neptune[3:] // "tune"
	println(tune)
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea"}
	dwarfSlice := []string{"Ceres", "Pluto", "Haumea"} // 直接创建切片
	fmt.Printf("%T %T", dwarfArray, dwarfSlice)        // [3]string []string
}
