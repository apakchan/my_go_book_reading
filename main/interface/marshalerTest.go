package main

import (
	"encoding/json"
	"fmt"
)

type newCoordinate struct {
	d, m, s float64
	h       rune
}

type decorateCoordinate struct {
	Decimal    float64 `json:"decimal"`
	Dms        string  `json:"dms"`
	Degrees    int     `json:"degrees"`
	Minutes    int     `json:"minutes"`
	Seconds    int     `json:"seconds"`
	Hemisphere string  `json:"hemisphere"`
}

func (coordinate newCoordinate) MarshalJSON() ([]byte, error) {
	resHandle := decorateCoordinate{
		Decimal:    coordinate.d + coordinate.m/60.0,
		Dms:        fmt.Sprintf("%.0f°%.0f'%.0f.0\" %v", coordinate.d, coordinate.m, coordinate.s, string(coordinate.h)),
		Degrees:    int(coordinate.d),
		Minutes:    int(coordinate.m),
		Seconds:    int(coordinate.s),
		Hemisphere: string(coordinate.h),
	}

	return json.Marshal(resHandle)
}

func main() {
	coordinate := newCoordinate{
		d: 154,
		m: 54,
		s: 0,
		h: 'E',
	}

	jsonByte, err := json.MarshalIndent(coordinate, "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonByte))
	}

}
