package solutions

/*

Given an array of integers, find and print the maximum number of integers you can select
from the array such that the absolute difference between any two of the chosen integers is
less than or equal to . For example, if your array is , you can create two subarrays
meeting the criterion:  and . The maximum length subarray has  elements.

Function Description

Complete the pickingNumbers function in the editor below. It should return an integer
that represents the length of the longest array that can be created.

pickingNumbers has the following parameter(s):

a: an array of integers

*/

// PickingNumbers returns the length of the longest array where all values are within 1 unit
func PickingNumbers(a []int32) int32 {
	var result []int32
	var maxlen int32 = 0
	for i, v := range a {
		result = append(result, v)
		for j, k := range a {
			if i != j {
				if result[0] >= k && absi(result[0]-k) <= 1 {
					result = append(result, k)
				}
			}
		}
		if int32(len(result)) > int32(len(a)) {
			return int32(len(result))
		}
		if int32(len(result)) > maxlen {
			maxlen = int32(len(result))
		}
		result = result[:0]
	}
	return maxlen
}

func absi(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
