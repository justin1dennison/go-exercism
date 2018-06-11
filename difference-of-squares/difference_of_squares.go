package diffsquares

func SquareOfSums(n int) int {
	total := 0
	for i := 0; i <= n; i += 1 {
		total += i
	}
	return square(total)
}

func SumOfSquares(n int) int {
	total := 0
	for i := 0; i <= n; i += 1 {
		total += square(i)
	}
	return total
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func square(n int) int {
	return n * n
}
