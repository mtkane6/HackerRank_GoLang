package solutions

/*

Karl has an array of integers. He wants to reduce the array until all remaining elements are equal. Determine the minimum number of elements to delete to reach his goal.

For example, if his array is , we see that he can delete the  elements  and  leaving . He could also delete both twos and either the  or the , but that would take  deletions. The minimum number of deletions is .

Function Description

Complete the equalizeArray function in the editor below. It must return an integer that denotes the minimum number of deletions required.

equalizeArray has the following parameter(s):

arr: an array of integers

*/

// EqualizeArray returns the number of items that are NOT the most common item
func EqualizeArray(arr []int32) int32 {
	var currentHigh int32 = 0
	var countMap = make(map[int32]int32)

	for _, v := range arr {
		countMap[v]++
		if countMap[v] > currentHigh {
			currentHigh = countMap[v]
		}
	}
	return int32(len(arr)) - currentHigh
}
