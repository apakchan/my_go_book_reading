package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

// 随机向一个空的存钱罐存 5 美分, 10 美分, 25 美分, 直到存款超过 20 美元位置
// 解决机械误差最好的方法就是不用浮点数
// 直接用美分为最小单位

func fmtPrintCurrentSum(curInput int, curSum int) {
	fmt.Printf("当前投入金额: %02d 美分, 当前总金额: %02d.%02d $\n", curInput, curSum/100, curSum%100)
}

func main() {
	dollarSlice := []int{5, 10, 25}
	sum := 0
	for sum <= 20*100 {
		curDollar := dollarSlice[rand.Intn(len(dollarSlice))]
		sum += curDollar
		fmtPrintCurrentSum(curDollar, sum)
	}

	const distance = 236000000000000000
	var lightSpeedTravelPerYear = big.NewInt(3 * 1e8 * 60 * 60 * 24 * 365)
	var distanceVar = big.NewInt(236000000000000000)
	var year = new(big.Int)
	fmt.Println(year.Div(distanceVar, lightSpeedTravelPerYear))
	fmt.Println(byte('A'))
}
