package solutions

import "fmt"

/*
Sam's house has an apple tree and an orange tree that yield an abundance of fruit. In the diagram below,
the red region denotes his house, where  is the start point, and  is the endpoint. The apple tree is to
the left of his house, and the orange tree is to its right. You can assume the trees are located on a
single point, where the apple tree is at point , and the orange tree is at point .


When a fruit falls from its tree, it lands  units of distance from its tree of origin along the -axis. A
negative value of  means the fruit fell  units to the tree's left, and a positive value of  means it falls
units to the tree's right.

Given the value of  for  apples and  oranges, determine how many apples and oranges will fall on Sam's house
(i.e., in the inclusive range )?

For example, Sam's house is between  and . The apple tree is located at  and the orange at . There are
pples and  oranges. Apples are thrown  units distance from , and  units distance. Adding each apple distance
to the position of the tree, they land at . Oranges land at . One apple and two oranges land in the inclusive
range so we print:

1
2

*/

// CountApplesAndOranges prints the number of apples and oranges that fall on the house
func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	countApples := 0
	for i := range apples {
		if apples[i]+a >= s && apples[i]+a <= t {
			countApples++
		}
	}
	countOranges := 0
	for i := range oranges {
		if oranges[i]+b >= s && oranges[i]+b <= t {
			countOranges++
		}
	}
	fmt.Println(countApples)
	fmt.Println(countOranges)

}
