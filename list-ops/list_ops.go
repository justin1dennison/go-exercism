package listops

type IntList []int
type predFunc func(int) bool
type binFunc func(int, int) int
type unaryFunc func(int) int

func (list IntList) Foldl(fn binFunc, initial int) int {
	if list.Length() == 0 {
		return initial
	}
	x, xs := list[0], list[1:]
	result := fn(x, initial)
	for _, el := range xs {
		result = fn(result, el)
	}
	return result
}
func (list IntList) Foldr(fn binFunc, initial int) int {
	if list.Length() == 0 {
		return initial
	}
	x, xs := list[0], list[1:]
	result := fn(initial, x)
	for _, el := range xs {
		result = fn(el, result)
	}
	return result
}
func (list IntList) Filter(fn predFunc) IntList {
	results := make([]int, 0)
	for _, el := range list {
		if fn(el) {
			results = append(results, el)
		}
	}
	return results
}
func (list IntList) Length() int {
	total := 0
	for i := 0; i < len(list); i++ {
		total += 1
	}
	return total
}
func (list IntList) Map(fn unaryFunc) IntList {
	results := make([]int, 0)
	for _, el := range list {
		results = append(results, fn(el))
	}
	return results
}
func (list IntList) Reverse() IntList {
	last := list.Length() - 1
	results := make([]int, 0)
	for i := last; i >= 0; i-- {
		results = append(results, list[i])
	}
	return results
}
func (list IntList) Append(lst IntList) IntList {
	results := make([]int, 0)
	for _, el := range list {
		results = append(results, el)
	}
	for _, el := range lst {
		results = append(results, el)
	}
	return results
}
func (list IntList) Concat(lists []IntList) IntList {
	results := make([]int, 0)
	for _, el := range list {
		results = append(results, el)
	}
	for _, lst := range lists {
		for _, el := range lst {
			results = append(results, el)
		}
	}
	return results
}
