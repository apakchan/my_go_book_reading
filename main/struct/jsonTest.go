package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Pos string  `json:"position"`
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitud"`
}

func main() {
	curiosity := location{
		Pos: "point1",
		Lat: -4.123,
		Lon: 123.123,
	}

	// convert to byte[]
	bytes, err := json.Marshal(curiosity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	indent, _ := json.MarshalIndent(curiosity, "", "")
	fmt.Println(string(bytes))
	fmt.Println(string(indent))
}
