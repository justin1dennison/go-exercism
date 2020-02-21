package lsproduct

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func chunks(input string, length int) []string {
	var cs []string
	for i := 0; i <= len(input)-length; i++ {
		chunk := input[i : i+length]
		cs = append(cs, chunk)
	}
	return cs
}

func digits(n string) ([]int, error) {
	var ds []int
	pieces := strings.Split(n, "")
	for _, piece := range pieces {
		converted, err := strconv.Atoi(piece)
		if err != nil {
			return ds, err
		}
		ds = append(ds, converted)
	}
	return ds, nil
}

func product(ns []int) int {
	start := 1
	for _, n := range ns {
		start *= n
	}
	return start
}

// LargestSeriesProduct determines the largest product from a contiguous set of digits
func LargestSeriesProduct(input string, length int) (int, error) {
	if length < 0 {
		return 1, errors.New("span must be greater than zero")
	}
	if len(input) < length {
		return -1, errors.New("span must be smaller than string length")
	}
	max := math.MinInt32
	cs := chunks(input, length)
	for _, c := range cs {
		ds, err := digits(c)
		if err != nil {
			return max, err
		}
		pd := product(ds)
		if pd > max {
			max = pd
		}
	}
	return max, nil
}
