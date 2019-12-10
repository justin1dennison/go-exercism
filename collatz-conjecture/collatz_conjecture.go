package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Out of range")
	}
	if n == 1 {
		return 0, nil
	}
	count := 0
	for n != 1 {
		n = step(n)
		count++
	}
	return count, nil
}

func step(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
