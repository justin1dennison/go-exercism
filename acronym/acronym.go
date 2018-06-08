package acronym

import (
	_ "fmt"
	"strings"
)

func Abbreviate(s string) string {
	temp := strings.Split(s, "-")
	words := make([]string, 0)
	firstChars := make([]string, len(words))
	//gather all of the words after both splits
	for _, t := range temp {
		chunks := strings.Split(t, " ")
		for _, chunk := range chunks {
			words = append(words, chunk)
		}
	}
	//gather all of the first characters as capitals
	for _, word := range words {
		first := strings.ToUpper(string(word[0]))
		firstChars = append(firstChars, first)
	}
	return strings.Join(firstChars, "")
}
