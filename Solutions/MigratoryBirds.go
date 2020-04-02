package solutions

/*

You have been asked to help study the population of birds migrating across the continent.
Each type of bird you are interested in will be identified by an integer value. Each time
a particular kind of bird is spotted, its id number will be added to your array of sightings.
You would like to be able to find out which type of bird is most common given a list of
sightings. Your task is to print the type number of that bird and if two or more types of
birds are equally common, choose the type with the smallest ID number.

For example, assume your bird sightings are of types . There are two each of types  and ,
and one sighting of type . Pick the lower of the two types seen twice: type .

*/

// MigratoryBirds returns the most commonly sighted bird with the lowest identifier
func MigratoryBirds(arr []int32) int32 {
	birdmap := make(map[int32]int)
	for _, v := range arr {
		birdmap[v]++
	}

	count := 0
	var maxBird int32 = 0

	for k, v := range birdmap {
		if v > count {
			maxBird = k
			count = v
		}
		if v == count && k < maxBird {
			maxBird = k
		}
	}
	return maxBird
}
