package solutions

import "math"

/*

Watson likes to challenge Sherlock's math ability. He will provide a starting and ending value describing a range of integers. Sherlock must determine the number of square integers within that range, inclusive of the endpoints.

Note: A square integer is an integer which is the square of an integer, e.g. .

For example, the range is  and , inclusive. There are three square integers in the range:  and .

Function Description

Complete the squares function in the editor below. It should return an integer representing the number of square integers in the inclusive range from  to .

squares has the following parameter(s):

a: an integer, the lower range boundary
b: an integer, the uppere range boundary

*/

// Squares returns the number of values with a whole number square root between a and b
func Squares(a int32, b int32) int32 {
	return int32(math.Floor(math.Sqrt(float64(b))) - math.Ceil(math.Sqrt(float64(a))) + 1)
}
