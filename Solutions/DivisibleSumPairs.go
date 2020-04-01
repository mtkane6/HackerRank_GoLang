package solutions

/*

You are given an array of  integers, , and a positive integer, . Find and print the number of  pairs where  and  +  is divisible by .

For example,  and . Our three pairs meeting the criteria are  and .

Function Description

Complete the divisibleSumPairs function in the editor below. It should return the integer count of pairs meeting the criteria.

divisibleSumPairs has the following parameter(s):

n: the integer length of array
ar: an array of integers
k: the integer to divide the pair sum by

*/

// DivisibleSumPairs returns the number of index combos whos sum is divisable by 'k'
func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	pairs := 0

	// 1 3 2 6 1 2
	for i := range ar {
		for j := range ar[i+1:] {
			if (ar[i]+ar[j+i+1])%k == 0 {
				pairs++
			}
		}
	}
	return int32(pairs)
}
