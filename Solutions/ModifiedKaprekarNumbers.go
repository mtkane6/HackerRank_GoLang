package solutions

import (
	"fmt"
	"strconv"
)

/*

A modified Kaprekar number is a positive whole number with a special property. If you square it, then split the number into two integers and sum those integers, you have the same value you started with.

Consider a positive whole number  with  digits. We square  to arrive at a number that is either  digits long or  digits long. Split the string representation of the square into two parts,  and . The right hand part,  must be  digits long. The left is the remaining substring. Convert those two substrings back to integers, add them and see if you get .

For example, if ,  then . We split that into two strings and convert them back to integers  and . We test , so this is not a modified Kaprekar number. If , still , and . This gives us , the original .

Note: r may have leading zeros.

Here's an explanation from Wikipedia about the ORIGINAL Kaprekar Number (spot the difference!):

In mathematics, a Kaprekar number for a given base is a non-negative integer, the representation of whose square in that base can be split into two parts that add up to the original number again. For instance, 45 is a Kaprekar number, because 45² = 2025 and 20+25 = 45.

Given two positive integers  and  where  is lower than , write a program to print the modified Kaprekar numbers in the range between  and , inclusive.

Function Description

Complete the kaprekarNumbers function in the editor below. It should print the list of modified Kaprekar numbers in ascending order.

kaprekarNumbers has the following parameter(s):

p: an integer
q: an integer

*/

// KaprekarNumbers prints all the Kaprekar Numberes within the range
func KaprekarNumbers(p int32, q int32) {

	var i, count int64 = int64(p), 0
	for ; i <= int64(q); i++ {
		var pSquared, leftInt, rightInt int64 = i * i, 0, 0
		numStr := fmt.Sprint(pSquared)
		leftStr := numStr[:len(numStr)/2]
		rightStr := numStr[len(numStr)/2:]

		var leftNum, rightNum int

		leftNum, _ = strconv.Atoi(leftStr)
		leftInt = int64(leftNum)
		rightNum, _ = strconv.Atoi(rightStr)
		rightInt = int64(rightNum)

		if leftInt+rightInt == i {
			fmt.Print(i, " ")
			count++
		}
	}
	if count == 0 {
		fmt.Println("INVALID RANGE")
	}
}
