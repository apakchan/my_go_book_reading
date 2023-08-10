package main

type kelvin float64
type celsius float64

// 函数
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// 方法
// (k kelvin) : 类型
// celsius()  : 方法名
// celsius    : 返回值
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	c := kelvin(275).celsius()
	c2 := kelvin(c)
	println(c)
	println(c2)
}
