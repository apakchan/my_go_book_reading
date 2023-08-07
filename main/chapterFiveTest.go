package main

import (
	"fmt"
	"math/rand"
)

type Ticket struct {
	company  string
	days     int
	flyTypes string
	price    int
}

const DISTANCE = 62100000

func main() {
	companyList := []string{"Space Adventures", "Space X", "Virgin Galactic"}
	fmtPrintTitle()
	for i := 0; i < 10; i++ {
		speed := rand.Intn(30-15) + 16 // [16, 31) km/s
		flyType := rand.Intn(2) == 1
		ticket := Ticket{
			company: companyList[rand.Intn(3)],
			days:    DISTANCE / speed / 60 / 60 / 24,
		}
		if flyType {
			ticket.flyTypes = "单程"
			ticket.price = getMoneyBySpeed(speed)
		} else {
			ticket.flyTypes = "往返"
			ticket.price = getMoneyBySpeed(speed) * 2
		}
		fmtPrintTicket(&ticket)
	}
}

func fmtPrintTitle() {
	fmt.Printf("%-15s %13s %12s %13s\n", "太空航空公司", "飞行天数", "飞行类型", "价格(百万美元)")
}

func fmtPrintTicket(ticket *Ticket) {
	fmt.Printf("%-20s %15d %15s %15d\n", ticket.company, ticket.days, ticket.flyTypes, ticket.price)
}

func getMoneyBySpeed(speed int) int {
	return 100*speed + 2000
}
