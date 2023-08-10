package main

// 这里会将数组整个复制传入函数中
// 注意数组长度也是数组的一部分
// 以 [5]string 作为入参和 [8]string 作为入参的函数是不同的
// 因此一般用 切片 作为函数入参, 而不是数组
func teraform(planets [...]string) {
	for _, ele := range planets {
		println("New " + ele)
	}
}

func main() {
	// 初始化数组默认赋值空字符串
	var planets1 [8]string
	for _, ele := range planets1 {
		println(ele)
	}
	planets2 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	for _, ele := range planets2 {
		println(ele)
	}
	// 复制数组
	planets3 := planets2
	planets3[2] = "myPlanet"
	println(planets2[2])
	println(planets3[2])
}
