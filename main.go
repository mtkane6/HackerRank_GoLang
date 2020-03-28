package main

import (
	s "./Solutions"
)

func main() {
	var n int32 = 6
	s.Staircase(n)

	x := []int32{1, 3, 5, 7, 9}
	s.MiniMaxSum(x)
}
