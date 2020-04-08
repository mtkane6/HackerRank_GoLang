package solutions

/*

Lily likes to play games with integers. She has created a new game where she determines the difference between a number and its reverse. For instance, given the number , its reverse is . Their difference is . The number  reversed is , and their difference is .

She decides to apply her game to decision making. She will look at a numbered range of days and will only go to a movie on a beautiful day.

Given a range of numbered days,  and a number , determine the number of days in the range that are beautiful. Beautiful numbers are defined as numbers where  is evenly divisible by . If a day's value is a beautiful number, it is a beautiful day. Print the number of beautiful days in the range.

Function Description

Complete the beautifulDays function in the editor below. It must return the number of beautiful days in the range.

beautifulDays has the following parameter(s):

i: the starting day number
j: the ending day number
k: the divisor

*/

// BeautifulDays returns the number of beautiful days between i and j inclusive
func BeautifulDays(i int32, j int32, k int32) int32 {
	var reversedNum int32 = 0
	var returnValue int32 = 0
	for i <= j {
		reversedNum = reverse(i, reversedNum)
		if abs(i-reversedNum)%k == 0 {
			returnValue++
		}
		reversedNum = 0
		i++
	}
	return returnValue
}

func reverse(x, result int32) int32 {
	for x > 0 {
		i := x % 10
		result = (result * 10) + i
		x = x / 10
	}
	return result
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
