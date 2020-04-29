package solutions

import "strconv"

/*

You are the benevolent ruler of Rankhacker Castle, and today you're distributing bread. Your subjects are in a line, and some of them already have some loaves. Times are hard and your castle's food stocks are dwindling, so you must distribute as few loaves as possible according to the following rules:

Every time you give a loaf of bread to some person , you must also give a loaf of bread to the person immediately in front of or behind them in the line (i.e., persons  or ).
After all the bread is distributed, each person must have an even number of loaves.
Given the number of loaves already held by each citizen, find and print the minimum number of loaves you must distribute to satisfy the two rules above. If this is not possible, print NO.

For example, the people in line have loaves . We can first give a loaf to  and  so . Next we give a loaf to  and  and have  which satisfies our conditions. We had to distribute  loaves.

Function Description

Complete the fairRations function in the editor below. It should return an integer that represents the minimum number of loaves required.

fairRations has the following parameter(s):

B: an array of integers that represent the number of loaves each persons starts with .

*/

// FairRations returns the minimum amount of bread loaves to give out
func FairRations(B []int32) string {
	var sum, count int32 = 0, 0
	for _, v := range B {
		sum += v
	}
	if sum%2 == 1 {
		return "NO"
	}

	for i := 0; i < len(B); i++ {
		if i == len(B)-1 && B[i]%2 == 1 {
			B[i]++
			B[i-1]++
			count += 2
		} else if B[i]%2 == 1 {
			B[i]++
			B[i+1]++
			count += 2
		}
	}
	return strconv.Itoa(int(count))
}
