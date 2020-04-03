package solutions

import "fmt"

/*

Two cats and a mouse are at various positions on a line. You will be given their starting positions. Your task is to determine which cat will reach the mouse first, assuming the mouse doesn't move and the cats travel at equal speed. If the cats arrive at the same time, the mouse will be allowed to move and it will escape while they fight.

You are given  queries in the form of , , and  representing the respective positions for cats  and , and for mouse . Complete the function  to return the appropriate answer to each query, which will be printed on a new line.

If cat  catches the mouse first, print Cat A.
If cat  catches the mouse first, print Cat B.
If both cats reach the mouse at the same time, print Mouse C as the two cats fight and mouse escapes.
For example, cat  is at position  and cat  is at . If mouse  is at position , it is  units from cat  and  unit from cat . Cat  will catch the mouse.

Function Description

Complete the catAndMouse function in the editor below. It should return one of the three strings as described.

catAndMouse has the following parameter(s):

x: an integer, Cat 's position
y: an integer, Cat 's position
z: an integer, Mouse 's position

*/

// CatAndMouse determines which cat, if any, will catch the mouse
func CatAndMouse(x int32, y int32, z int32) string {
	if abs(x-z) < abs(y-z) {
		return fmt.Sprintf("Cat A")
	} else if abs(x-z) > abs(y-z) {
		return fmt.Sprintf("Cat B")
	} else {
		return fmt.Sprintf("Mouse C")
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
