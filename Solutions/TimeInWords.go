package solutions

// TimeInWords turns time in integers into time in words
func TimeInWords(h int32, m int32) string {
	timeWords := []string{
		"o' clock",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"quarter",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
		"twenty",
		"twenty one",
		"twenty two",
		"twenty three",
		"twenty four",
		"twenty five",
		"twenty six",
		"twenty seven",
		"twenty eight",
		"twenty nine",
		"half past ",
	}

	if m == 0 || m == 00 {
		return timeWords[h] + " " + timeWords[0]
	}
	if m == 1 || m == 01 {
		return timeWords[m] + " minute past " + timeWords[h]
	}
	if m == 15 {
		return timeWords[m] + " past " + timeWords[h]
	}
	if m == 30 {
		return timeWords[30] + timeWords[h]
	}
	if m == 45 {
		return timeWords[15] + " to " + timeWords[(h+1)%12]
	}
	if m > 30 {
		if m == 59 {
			return timeWords[60-m] + " minute to " + timeWords[(h+1)%12]
		}
		return timeWords[60-m] + " minutes to " + timeWords[(h+1)%12]
	}
	return timeWords[m] + " minutes past " + timeWords[h]
}
