package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	// 创建新的切片, 而不是修改原来的
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}

func main() {
	dwarfs := make([]string, 1, 10)               // type, len, cap
	fmt.Println(cap(dwarfs), len(dwarfs), dwarfs) // 10, 1
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea")
	fmt.Println(cap(dwarfs), len(dwarfs), dwarfs) // 第一个是零值元素

	fmt.Println(terraform("prefix", dwarfs...))

}
