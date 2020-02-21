package wordcount

import (
	"regexp"
	"strings"
)

var (
	splitter = regexp.MustCompile(` |\n`)
	punc     = regexp.MustCompile("[^a-zA-Z0-9']+")
)

// Frequency keeps track of occurrence of words
type Frequency map[string]int

// WordCount calculates the word frequency
func WordCount(input string) Frequency {
	fs := make(Frequency)
	cleaned := normalize(input)
	words := splitter.Split(cleaned, -1)
	for _, word := range words {
		if word != " " {
			t := strings.Trim(word, "'")
			fs[t]++
		}
	}
	return fs
}

func normalize(input string) string {
	removed := punc.ReplaceAllString(input, " ")
	trimmed := strings.Trim(removed, " ")
	return strings.ToLower(trimmed)
}
