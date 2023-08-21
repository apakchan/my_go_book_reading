package main

import "fmt"

func loopAddEle(srcSlice []string, eles ...string) {
	startCap := cap(srcSlice)
	for _, s := range eles {
		curCap := cap(srcSlice)
		capAddInThisLoop := cap(srcSlice) == len(srcSlice)
		if capAddInThisLoop {
			fmt.Printf("now capacity: %d\n", curCap)
		}
		srcSlice = append(srcSlice, s)
		if capAddInThisLoop {
			fmt.Printf("from %d to %d\n", curCap, cap(srcSlice))
		}
	}
	endCap := cap(srcSlice)
	if startCap != 0 {
		fmt.Printf("容量增加至 %d 倍\n", endCap/startCap)
	} else {
		fmt.Printf("容量从 0 增加到 %d\n", endCap)
	}
}

func main() {
	eles := make([]string, 0)
	for i := 0; i < 100; i++ {
		eles = append(eles, "i")
	}
	loopAddEle(make([]string, 1), eles...)
}
