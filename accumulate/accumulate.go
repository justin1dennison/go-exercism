package accumulate


func Accumulate(input []string, fn func(string) string) []string {
	var r []string
	for _, in := range input {
		r = append(r, fn(in))
	}
	return r
}
