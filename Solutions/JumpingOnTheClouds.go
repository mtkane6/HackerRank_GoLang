package solutions

/*

Aerith is playing a cloud hopping game. In this game, there are sequentially numbered clouds that can be thunderheads or cumulus clouds. Her character must jump from cloud to cloud until it reaches the start again.

To play, Aerith is given an array of clouds,  and an energy level . She starts from  and uses  unit of energy to make a jump of size  to cloud . If Aerith lands on a thundercloud, , her energy () decreases by  additional units. The game ends when Aerith lands back on cloud .

Given the values of , , and the configuration of the clouds as an array , can you determine the final value of  after the game ends?

For example, give  and , the indices of her path are . Her energy level reduces by  for each jump to . She landed on one thunderhead at an additional cost of  energy units. Her final energy level is .

Note: Recall that  refers to the modulo operation. In this case, it serves to make the route circular. If Aerith is at  and jumps , she will arrive at .

Function Description

Complete the jumpingOnClouds function in the editor below. It should return an integer representing the energy level remaining after the game.

jumpingOnClouds has the following parameter(s):

c: an array of integers representing cloud types
k: an integer representing the length of one jump

*/

// JumpingOnClouds returns the amount of remaining jumping power
func JumpingOnClouds(c []int32, k int32) int32 {
	var power int32 = 100
	var i int32 = 0

	for {
		i = (i + k) % int32(len(c))
		power--
		if c[i] == 1 {
			power -= 2
		}
		if i == 0 {
			return power
		}
	}
}
