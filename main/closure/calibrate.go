package main

import "math/rand"

type kelvin float64

// sensor 是一个函数类型的变量
type sensor func() kelvin

func realSensor() kelvin {
	return kelvin(150 + rand.Intn(151))
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	i1 := kelvin(1)
	s1 := calibrate(func() kelvin {
		return 1
	}, i1)()
	println(s1)
	i1++
	s2 := calibrate(func() kelvin {
		return 1
	}, i1)()
	// 仍然是 2
	// 因为 offset 副本接收的是实参的副本值而不是引用
	println(s1)
	println(s2)

	s3 := realSensor
	println(calibrate(s3, 1)())
	println(calibrate(s3, 1)())
	println(calibrate(s3, 1)())
	println(calibrate(s3, 1)())
}
