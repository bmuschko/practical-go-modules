package main

import (
	"fmt"
	"github.com/bmuschko/calc"
	calcv2 "github.com/bmuschko/calc/v2"
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calcv2.Add(1, 2, 5, 8))
}
