package solutions

import (
	"strconv"
)

// TimeConversion converts 12-hour time to 24-hour time
func TimeConversion(s string) string {
	if s[len(s)-2] == byte('P') || s[len(s)-2] == byte('p') {
		if s[0] == '1' && s[1] == '2' {
			result := s[:8]
			return result
		}
		hour, _ := (strconv.Atoi(s[:2]))
		intHour := (hour + 12) % 24
		result := strconv.Itoa(intHour) + s[2:]
		return result[:8]
	} else if s[1] == '2' {
		result := "00" + s[2:8]
		return result
	} else {
		return s[:8]
	}
}
