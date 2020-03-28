package solutions

// BirthdayCakeCandles return the number of tallest birthday candles
func BirthdayCakeCandles(ar []int32) int32 {
	var max, count int32 = 0, 0

	for _, v := range ar {
		if v == max {
			count++
		}
		if v > max {
			max = v
			count = 1
		}
	}
	return count
}
