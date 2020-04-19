package solutions

/*

Given a sequence of integers , a triplet  is beautiful if:

Given an increasing sequenc of integers and the value of , count the number of beautiful triplets in the sequence.

For example, the sequence  and . There are three beautiful triplets, by index: . To test the first triplet,  and .

Function Description

Complete the beautifulTriplets function in the editor below. It must return an integer that represents the number of beautiful triplets in the sequence.

beautifulTriplets has the following parameters:

d: an integer
arr: an array of integers, sorted ascending

*/

// BeautifulTriplets returns the number of beautiful triplets per the description above
func BeautifulTriplets(d int32, arr []int32) int32 {
	var triplets int32 = 0
	foundTriplet := false
	for i := 0; i < len(arr)-2; i++ {
		foundTriplet = false
		j := i + 1
		for arr[j] <= arr[i]+d {
			if arr[j] == arr[i]+d {
				if j+1 < len(arr) {
					k := j + 1
					for arr[k] <= arr[j]+d {
						if arr[k] == arr[j]+d {
							triplets++
							foundTriplet = true
							break
						}
						k++
						if k == len(arr) {
							break
						}
					}
				}

			}
			if foundTriplet {
				break
			}
			j++
			if j == len(arr)-1 {
				break
			}
		}
	}
	return triplets
}
