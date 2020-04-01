package solutions

/*

Lily has a chocolate bar that she wants to share it with Ron for his birthday. Each of the squares has an integer on it. She decides to share a contiguous segment of the bar selected such that the length of the segment matches Ron's birth month and the sum of the integers on the squares is equal to his birth day. You must determine how many ways she can divide the chocolate.

Consider the chocolate bar as an array of squares, . She wants to find segments summing to Ron's birth day,  with a length equalling his birth month, . In this case, there are two segments meeting her criteria:  and .

Function Description

Complete the birthday function in the editor below. It should return an integer denoting the number of ways Lily can divide the chocolate bar.

birthday has the following parameter(s):

s: an array of integers, the numbers on each of the squares of chocolate
d: an integer, Ron's birth day
m: an integer, Ron's birth month
Input Format

The first line contains an integer , the number of squares in the chocolate bar.
The second line contains  space-separated integers , the numbers on the chocolate squares where .
The third line contains two space-separated integers,  and , Ron's birth day and his birth month.

*/

// Birthday returns the number of ways to provide 'm' chocolate squares that == 'd'.
func Birthday(s []int32, d int32, m int32) int32 {
	possibleWays := 0
	newM := int(m)
	for i := range s {
		if newM+i <= len(s) {
			currSquares := s[0+i : newM+i]
			if sum(currSquares) == d {
				possibleWays++
			}
		} else {
			break
		}
	}
	return int32(possibleWays)
}

func sum(squares []int32) int32 {
	var sum int32 = 0
	for _, v := range squares {
		sum += v
	}
	return sum
}
