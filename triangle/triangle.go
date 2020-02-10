package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func isValidTriangle(a, b, c float64) bool {
	positiveSides := a > 0 && b > 0 && c > 0
	finiteSides := !math.IsInf(a, 0) && !math.IsInf(b, 0) && !math.IsInf(c, 0)
	satisfiesInequality := (a+b >= c) && (a+c >= b) && (b+c >= a)
	return positiveSides && satisfiesInequality && finiteSides
}

// KindFromSides returns the type of triangle based on the sides
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c && a == c {
		return Equ
	} else if (a == b && b != c) || (a == c && b != c) || (b == c && a != c) {
		return Iso
	} else {
		return Sca
	}
}
