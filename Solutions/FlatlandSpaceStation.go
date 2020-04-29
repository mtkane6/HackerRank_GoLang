package solutions

/*

Flatland is a country with a number of cities, some of which have space stations. Cities are numbered consecutively and each has a road of  length connecting it to the next city. It is not a circular route, so the first city doesn't connect with the last city. Determine the maximum distance from any city to it's nearest space station.

For example, there are  cities and  of them has a space station, city . They occur consecutively along a route. City  is  unit away and city  is  units away. City  is  units from its nearest space station as one is located there. The maximum distance is .

Function Description

Complete the flatlandSpaceStations function in the editor below. It should return an integer that represents the maximum distance any city is from a space station.

flatlandSpaceStations has the following parameter(s):

n: the number of cities
c: an integer array that contains the indices of cities with a space station, -based indexing

*/

// FlatlandSpaceStations returns the longest distance from any city to the nearest space station
func FlatlandSpaceStations(n int32, c []int32) int32 {
	var spaceMap = make(map[int32]int32)
	for _, v := range c {
		spaceMap[v] = v
	}

	var i, maxDist int32 = 0, 0

	for ; i < n; i++ {
		if _, ok := spaceMap[i]; ok {
			continue
		}
		var nearest int32 = n
		for _, v := range c {
			currDist := abs(i, v)
			if currDist < nearest {
				nearest = currDist
			}
		}
		if nearest > maxDist {
			maxDist = nearest
		}
	}
	return maxDist
}

func abs(a, b int32) int32 {
	x := a - b
	if x < 0 {
		return -x
	}
	return x
}
