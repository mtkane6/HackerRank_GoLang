package solutions

import "fmt"

/*

You are given a number of sticks of varying lengths. You will iteratively cut the sticks into smaller sticks, discarding the shortest pieces until there are none left. At each iteration you will determine the length of the shortest stick remaining, cut that length from each of the longer sticks and then discard all the pieces of that shortest length. When all the remaining sticks are the same length, they cannot be shortened so discard them.

Given the lengths of  sticks, print the number of sticks that are left before each iteration until there are none left.

For example, there are  sticks of lengths . The shortest stick length is , so we cut that length from the longer two and discard the pieces of length . Now our lengths are . Again, the shortest stick is of length , so we cut that amount from the longer stick and discard those pieces. There is only one stick left, , so we discard that stick. Our lengths are .

Function Description

Complete the cutTheSticks function in the editor below. It should return an array of integers representing the number of sticks before each cut operation is performed.

cutTheSticks has the following parameter(s):

arr: an array of integers representing the length of each stick

*/

// CutTheSticks returns a slice of the remaining number of sticks after each round of cuts
func CutTheSticks(arr []int32) []int32 {
	var currentMin int32 = arr[0]
	var sticksRemaining []int32
	sticksRemaining = append(sticksRemaining, int32(len(arr)))

	currentMin = findMin(arr, currentMin)

	for len(arr) > 0 {
		arr = cut(arr, currentMin)
		fmt.Println("arr:", arr)
		if len(arr) > 0 {
			sticksRemaining = append(sticksRemaining, int32(len(arr)))
			currentMin = findMin(arr, arr[0])
		}
	}
	return sticksRemaining
}

func findMin(arr []int32, min int32) int32 {
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func cut(arr []int32, min int32) []int32 {
	var newArr []int32
	for _, v := range arr {
		if v > min {
			newArr = append(newArr, v-min)
		}
	}
	return newArr
}
