package robotname

import (
	"math/rand"
	"strings"
)

const nameLength int = 5
const letterLength int = 2
const digitLength int = nameLength - letterLength
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"

var seenNames map[string]bool

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}

	result := make([]string, nameLength)

	for i := 0; i < letterLength; i++ {
		r := getRandomChar(letters)
		result = append(result, r)
	}

	for i := 0; i < digitLength; i++ {
		d := getRandomChar(digits)
		result = append(result, d)
	}
	name := strings.Join(result, "")
	if _, ok := seenNames[name]; ok {
		name = r.Name()
	}
	r.name = name
	return name
}

func (r *Robot) Reset() {
	r.name = ""
}

func getRandomChar(source string) string {
	rs := strings.Split(source, "")
	i := rand.Intn(len(rs))
	return rs[i]
}
