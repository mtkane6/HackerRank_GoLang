package solutions

/*

Lilah has a string, , of lowercase English letters that she repeated infinitely many times.

Given an integer, , find and print the number of letter a's in the first  letters of Lilah's infinite string.

For example, if the string  and , the substring we consider is , the first  characters of her infinite string. There are  occurrences of a in the substring.

Function Description

Complete the repeatedString function in the editor below. It should return an integer representing the number of occurrences of a in the prefix of length  in the infinitely repeating string.

repeatedString has the following parameter(s):

s: a string to repeat
n: the number of characters to consider

*/

// RepeatedString returns the number of 'a's in the length of string
func RepeatedString(s string, n int64) int64 {
	var wholeAs, fractionAs, whole, fraction int64

	whole = n / int64(len(s))
	fraction = n % int64(len(s))

	for i, v := range s {
		if v == 'a' {
			wholeAs++
			if i < int(fraction) {
				fractionAs++
			}
		}
	}

	return (wholeAs * whole) + fractionAs
}
