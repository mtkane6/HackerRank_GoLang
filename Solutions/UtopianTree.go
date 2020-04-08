package solutions

/*

The Utopian Tree goes through 2 cycles of growth every year. Each spring, it doubles in height. Each summer, its height increases by 1 meter.

Laura plants a Utopian Tree sapling with a height of 1 meter at the onset of spring. How tall will her tree be after  growth cycles?

For example, if the number of growth cycles is , the calculations are as follows:

Period  Height
0          1
1          2
2          3
3          6
4          7
5          14
Function Description

Complete the utopianTree function in the editor below. It should return the integer height of the tree after the input number of growth cycles.

utopianTree has the following parameter(s):

n: an integer, the number of growth cycles to simulate

*/

// UtopianTree returns the height of the tree after n growth periods
func UtopianTree(n int32) int32 {
	newN := int(n)
	height := 0
	for i := 0; i <= newN; i++ {
		if i%2 == 0 {
			height++
		} else {
			height = double(height)
		}
	}
	return int32(height)
}

func double(x int) int {
	return x * 2
}
