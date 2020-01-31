package pangram

import "strings"

func IsPangram(word string) bool {
	letters := strings.Split("abcdefghijklmopqrstuvwxyz", "")
	lowercase := strings.ToLower(word)
	for _, letter := range letters {
		if strings.Index(lowercase, letter) == -1 {
			return false
		}
	}
	return true
}
