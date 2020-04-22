package solutions

/*

You wish to buy video games from the famous online video game store Mist.

Usually, all games are sold at the same price,  dollars. However, they are planning to have the seasonal Halloween Sale next month in which you can buy games at a cheaper price. Specifically, the first game you buy during the sale will be sold at  dollars, but every subsequent game you buy will be sold at exactly  dollars less than the cost of the previous one you bought. This will continue until the cost becomes less than or equal to  dollars, after which every game you buy will cost  dollars each.

For example, if ,  and , then the following are the costs of the first  games you buy, in order:

You have  dollars in your Mist wallet. How many games can you buy during the Halloween Sale?

*/

// HowManyGames returns how many games can be bought with your budget as the prices of games falls.
func HowManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var games, totalCost int32 = 0, 0

	for {
		totalCost += p
		if s >= totalCost {
			if p-d > m {
				p -= d
			} else {
				p = m
			}
			games++
		} else {
			return games
		}
	}
}
