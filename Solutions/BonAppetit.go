package solutions

import "fmt"

/*

Anna and Brian are sharing a meal at a restuarant and they agree to split the bill equally. Brian wants to order something that Anna is allergic to though, and they agree that Anna won't pay for that item. Brian gets the check and calculates Anna's portion. You must determine if his calculation is correct.

For example, assume the bill has the following prices: . Anna declines to eat item  which costs . If Brian calculates the bill correctly, Anna will pay . If he includes the cost of , he will calculate . In the second case, he should refund  to Anna.

Function Description

Complete the bonAppetit function in the editor below. It should print Bon Appetit if the bill is fairly split. Otherwise, it should print the integer amount of money that Brian owes Anna.

bonAppetit has the following parameter(s):

bill: an array of integers representing the cost of each item ordered
k: an integer representing the zero-based index of the item Anna doesn't eat
b: the amount of money that Anna contributed to the bill

*/

// BonAppetit returns whether or not Brian charged his gf correctly for food.
func BonAppetit(bill []int32, k int32, b int32) {
	newBill := bill[:k]
	newBill = append(newBill, bill[k+1:]...)

	var sum int32 = 0
	for _, v := range newBill {
		sum += v
	}

	if sum/2 == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - (sum / 2))
	}
}
