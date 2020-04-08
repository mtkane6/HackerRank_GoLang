package solutions

/*
Dan is playing a video game in which his character competes in a hurdle race. Hurdles are of varying heights, and Dan has a maximum height he can jump. There is a magic potion he can take that will increase his maximum height by  unit for each dose. How many doses of the potion must he take to be able to jump all of the hurdles.

Given an array of hurdle heights , and an initial maximum height Dan can jump, , determine the minimum number of doses Dan must take to be able to clear all the hurdles in the race.

For example, if  and Dan can jump  unit high naturally, he must take  doses of potion to be able to jump all of the hurdles.

Function Description

Complete the hurdleRace function in the editor below. It should return the minimum units of potion Dan needs to drink to jump all of the hurdles.

hurdleRace has the following parameter(s):

k: an integer denoting the height Dan can jump naturally
height: an array of integers denoting the heights of each hurdle

*/

// HurdleRace returns the amount of potion needed to jump over all the hurdles.
func HurdleRace(k int32, height []int32) int32 {
	var maxHeight int32 = 0
	for _, v := range height {
		if v > maxHeight {
			maxHeight = v
		}
	}
	if maxHeight > k {
		return maxHeight - k
	}
	return 0
}
