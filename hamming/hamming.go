package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The strings must be the same distance")
	}
	count := 0
	for index, first := range a {
		second := rune(b[index])
		if first != second {
			count += 1
		}
	}
	return count, nil

}
