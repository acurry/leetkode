package main

import (
	"fmt"
	"strconv"
	"strings"
)

func divide(dividend int, divisor int) int {
	val := float64(dividend / divisor)
	valAsStr := fmt.Sprintf("%f", val)
	whole := strings.Split(valAsStr, ".")[0]

	ret, _ := strconv.ParseInt(whole, 10, 0)

	return int(ret)
}

type args struct {
	Dividend int
	Divisor  int
}

func main() {
	data := []args{
		{
			Dividend: 10,
			Divisor:  3,
		},
		{
			Dividend: 10,
			Divisor:  -3,
		},
	}

	for _, a := range data {
		v := divide(a.Dividend, a.Divisor)
		fmt.Printf("divide(%d, %d): %d\n", a.Dividend, a.Divisor, v)
	}
	// fmt.Print("SOLVED!")
}
