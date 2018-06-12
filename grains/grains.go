package grains

import "errors"

func Square(n int) (uint64, error) {
	var k uint64 = 1
	if n < 1 || n > 64 {
		return 0, errors.New("Input must be between 1 and 64")
	}
	for i := 1; i < n; i += 1 {
		k *= 2
	}
	return k, nil
}

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i <= 64; i += 1 {
		val, _ := Square(i)
		total += val
	}
	return total
}
