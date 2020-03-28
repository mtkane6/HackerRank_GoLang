package solutions

import "fmt"

// func MiniMaxSum(arr []int32) {
// 	var min, max int32 = arr[0], arr[0]
// 	var minIndex, maxIndex int = 0, 0

// 	for i, v := range arr {
// 		if v < min {
// 			min = v
// 			minIndex = i
// 		}
// 		if v > max {
// 			max = v
// 			maxIndex = i
// 		}
// 	}
// 	var maxSlice = make([]int32, 4)
// 	var minSlice = make([]int32, 4)
// 	maxSlice = append(maxSlice, arr[:minIndex]...)
// 	maxSlice = append(maxSlice, arr[minIndex+1:]...)

// 	minSlice = append(minSlice, arr[:maxIndex]...)
// 	minSlice = append(minSlice, arr[maxIndex+1:]...)

// 	fmt.Println(minSlice)
// 	fmt.Println(maxSlice)

// 	var minResult, maxResult int32 = 0, 0

// 	for i := range minSlice {
// 		minResult += minSlice[i]
// 		maxResult += maxSlice[i]
// 	}
// 	fmt.Println(minResult, maxResult)

// }

// MiniMaxSum returns the sum of the smallest 4/5 and largest 4/5 values
func MiniMaxSum(arr []int32) {
	if len(arr) == 0 {
		fmt.Println(0, 0)
	}
	if len(arr) == 1 {
		fmt.Println(arr[0], arr[0])
	}

	var min, max int32 = arr[0], arr[0]
	var sum int64 = 0

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += int64(v)
	}

	fmt.Println(sum-int64(max), sum-int64(min))
}
