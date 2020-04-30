package solutions

/*

You are given a square map as a matrix of integer strings. Each cell of the map has a value denoting its depth. We will call a cell of the map a cavity if and only if this cell is not on the border of the map and each cell adjacent to it has strictly smaller depth. Two cells are adjacent if they have a common side, or edge.

Find all the cavities on the map and replace their depths with the uppercase character X.

For example, given a matrix:

989
191
111
You should return:

989
1X1
111
The center cell was deeper than those on its edges: [8,1,1,1]. The deep cells in the top two corners don't share an edge with the center cell.

Function Description

Complete the cavityMap function in the editor below. It should return an array of strings, each representing a line of the completed map.

cavityMap has the following parameter(s):

grid: an array of strings, each representing a row of the grid
Input Format

*/

// CavityMap returns all non-edge high values in the square matrix of strings
func CavityMap(grid []string) []string {

	var returnByteArr [][]byte

	for _, v := range grid {
		returnByteArr = append(returnByteArr, []byte(v))
	}
	var result []string
	result = append(result, string(returnByteArr[0]))

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] > grid[i-1][j] &&
				grid[i][j] > grid[i+1][j] &&
				grid[i][j] > grid[i][j-1] &&
				grid[i][j] > grid[i][j+1] {
				returnByteArr[i][j] = byte('X')
			}
		}
		result = append(result, string(returnByteArr[i]))
	}
	if len(returnByteArr) > 1 {
		result = append(result, string(returnByteArr[len(returnByteArr)-1]))
	}
	return result
}
