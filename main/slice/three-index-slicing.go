package main

import "fmt"

// 与其问什么时候用三索引切分, 不如问什么时候不用三索引切分
// 将整个底层数组切分下来进行 append 操作不会影响被切分数组
// 如果是切分前一部分数组进行任何写操作, 那么会影响被切分数组的
func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terrestrial := planets[0:4:4]          // 长度为 4, 容量为 4  如果是 planets[0:4], 容量为 8
	worlds := append(terrestrial, "Ceres") // double capacity, malloc different array
	fmt.Println(worlds)
}
