package leap

// IsLeapYear returns whether a provided year is a leap year
func IsLeapYear(n int) bool {
	return (n%4 == 0 && n%100 != 0) || n%400 == 0
}
