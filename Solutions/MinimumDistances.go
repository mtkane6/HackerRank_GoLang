package solutions

/*

We define the distance between two array values as the number of indices between the two values. Given , find the minimum distance between any pair of equal elements in the array. If no such value exists, print .

For example, if , there are two matching pairs of values: . The indices of the 's are  and , so their distance is . The indices of the 's are  and , so their distance is .

Function Description

Complete the minimumDistances function in the editor below. It should return the minimum distance between any two matching elements.

minimumDistances has the following parameter(s):

a: an array of integers

*/

// MinimumDistances returns the minimum distance between indexes of pairs of integer values
func MinimumDistances(a []int32) int32 {
	var distanceMap = make(map[int32]int)
	var minDistance int = len(a)

	for index, key := range a {
		if val, ok := distanceMap[key]; ok {
			dist := index - val
			if dist < minDistance {
				minDistance = dist
			}
		} else {
			distanceMap[key] = index
		}
	}
	if minDistance == len(a) {
		return int32(-1)
	}
	return int32(minDistance)
}
