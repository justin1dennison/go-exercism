package isogram

import "strings"

func IsIsogram(s string) bool {
	charsToCount := make(map[string]int)
	chars := strings.Split(s, "")
	for _, char := range chars {
		upperCase := strings.ToUpper(char)
		if _, ok := charsToCount[upperCase]; ok {
			charsToCount[upperCase] += 1
		} else {
			charsToCount[upperCase] = 1
		}
	}

	for key, val := range charsToCount {
		if key == "-" || key == " " {
			continue
		}
		if val > 1 {
			return false
		}
	}
	return true

}
