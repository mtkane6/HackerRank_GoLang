package solutions

/*

An integer  is a divisor of an integer  if the remainder of .

Given an integer, for each digit that makes up the integer determine whether it is a divisor. Count the number of divisors occurring within the integer.

Note: Each digit is considered to be unique, so each occurrence of the same digit should be counted (e.g. for ,  is a divisor of  each time it occurs so the answer is ).

Function Description

Complete the findDigits function in the editor below. It should return an integer representing the number of digits of  that are divisors of .

findDigits has the following parameter(s):

n: an integer to analyze

*/

// FindDigits returns the number of digits of an integer that are divisors of that integer
func FindDigits(n int32) int32 {
	var copy, divisor int32 = n, 0

	for copy != 0 {
		digit := copy % int32(10)
		if digit != 0 {
			if n%digit == 0 {
				divisor++
			}
		}
		copy = copy / 10
	}
	return divisor
}
