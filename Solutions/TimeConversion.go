package solutions

import (
	"strconv"
)

// TimeConversion converts 12-hour time to 24-hour time
func TimeConversion(s string) string {
	if s[len(s)-2] == byte('P') || s[len(s)-2] == byte('p') {
		if s[0] == '1' && s[1] == '2' {
			return s[:8]
		}
		hour, _ := (strconv.Atoi(s[:2]))
		return strconv.Itoa((hour+12)%24) + s[2:8]
	} else if s[1] == '2' {
		return "00" + s[2:8]
	} else {
		return s[:8]
	}
}
