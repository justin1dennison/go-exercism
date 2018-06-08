package twofer

import "fmt"

func ShareWith(s string) string {
	var name string
	if s == "" {
		name = "you"
	} else {
		name = s
	}
	phrase := fmt.Sprintf("One for %s, one for me.", name)
	return phrase
}
