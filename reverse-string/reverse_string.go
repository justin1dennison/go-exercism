package reverse

import "strings"

func String(s string) string {
	chars := strings.Split(s, "")
	result := ""
	for i := len(chars) - 1; i >= 0; i -= 1 {
		result += chars[i]
	}

	return result
}
