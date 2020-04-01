package main

import (
	"fmt"

	s "./Solutions"
)

func main() {
	var n int32 = 6
	s.Staircase(n)

	x := []int32{1, 3, 5, 7, 9}
	s.MiniMaxSum(x)

	time := "02:34:00PM"
	fmt.Println(s.TimeConversion(time))

	for i := range x[:] {
		for j := range x[i+1:] {
			fmt.Println(i, j+i+1)
		}
	}
}
