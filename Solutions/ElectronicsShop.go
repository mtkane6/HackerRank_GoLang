package solutions

/*

Monica wants to buy a keyboard and a USB drive from her favorite electronics store. The store has several models of each. Monica wants to spend as much as possible for the  items, given her budget.

Given the price lists for the store's keyboards and USB drives, and Monica's budget, find and print the amount of money Monica will spend. If she doesn't have enough money to both a keyboard and a USB drive, print -1 instead. She will buy only the two required items.

For example, suppose she has  to spend. Three types of keyboards cost . Two USB drives cost . She could purchase a , or a . She chooses the latter. She can't buy more than  items so she can't spend exactly .

Function Description

Complete the getMoneySpent function in the editor below. It should return the maximum total price for the two items within Monica's budget, or  if she cannot afford both items.

getMoneySpent has the following parameter(s):

keyboards: an array of integers representing keyboard prices
drives: an array of integers representing drive prices
b: the units of currency in Monica's budget

*/

// GetMoneySpent returns the highest priced combo of keyboard/drive possible
func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maxCost int32 = -1
	for _, k := range keyboards {
		for _, d := range drives {
			if k+d > maxCost && k+d <= b {
				maxCost = k + d
			}
		}
	}
	return maxCost
}
