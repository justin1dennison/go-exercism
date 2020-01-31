package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	rs := make(map[string]int)
	for k, xs := range input {
		for _, x := range xs {
			rs[strings.ToLower(x)] = k
		}
	}
	return rs
}
