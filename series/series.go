package series

// All returns all of the substrings of length n
func All(n int, s string) []string {
	var r []string
	for index := 0; index < len(s)-n+1; index++ {
		r = append(r, s[index:index+n])
	}
	return r
}

// UnsafeFirst returns the first substring of s that is n in length
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First returns the first substring of s for a given length n and whether that was successful or not
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
