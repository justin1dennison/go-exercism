package hamming

import "errors"

func Distance(a, b string) (int, error) {
	aChars := []rune(a)
	bChars := []rune(b)
	if len(aChars) != len(bChars) {
		return -1, errors.New("The strings must be the same distance")
	}
	count := 0
	for i, first := range aChars {
		second := bChars[i]
		if first != second {
			count += 1
		}
	}
	return count, nil

}
