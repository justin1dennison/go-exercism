package isbn

import (
	"regexp"
	"strings"
)

var (
	letterPattern = regexp.MustCompile("[A-WY-Z]")
)

//IsValidISBN validates whether the isbn is valid using the ISBN-10 verification process
func IsValidISBN(isbn string) bool {
	n := normalize(isbn)
	if !isCorrectLength(n) ||
		containsLetters(n) ||
		xIsNotCheckDigit(n) {
		return false
	}

	var v int
	for i, r := range n {
		var value int
		if r == 'X' {
			value = 10
		} else {
			value = int(r) - '0'
		}

		v += (value * (10 - i))
	}
	return v%11 == 0
}

func normalize(isbn string) string {
	return strings.ReplaceAll(isbn, "-", "")
}

func isCorrectLength(s string) bool {
	return len(s) == 10
}

func containsLetters(s string) bool {
	return letterPattern.Match([]byte(s))

}

func xIsNotCheckDigit(s string) bool {
	return strings.Contains(s, "X") && !strings.HasSuffix(s, "X")
}
