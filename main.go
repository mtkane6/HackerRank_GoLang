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
}
