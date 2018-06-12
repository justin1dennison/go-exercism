package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Valid(s string) bool {
	reg := regexp.MustCompile(`[^0-9\s]+`)
	stripped := strings.Replace(s, " ", "", -1)
	chars := strings.Split(stripped, "")
	if len(chars) <= 1 || reg.MatchString(s) {
		return false
	}
	digits := getDigits(chars)
	doubled := doubleEverySecondFromRight(digits)
	total := sum(doubled)
	return total%10 == 0
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func doubleEverySecondFromRight(nums []int) []int {
	result := make([]int, len(nums))
	isSecond := false
	for i := len(nums) - 1; i >= 0; i -= 1 {
		if !isSecond {
			result[i] = nums[i]
		} else {
			if nums[i] >= 5 {
				result[i] = (nums[i] * 2) - 9
			} else {
				result[i] = nums[i] * 2
			}
		}
		isSecond = !isSecond

	}
	return result
}

func getDigits(chars []string) []int {
	digits := make([]int, 0)
	for i := range chars {
		converted, err := strconv.Atoi(chars[i])
		if err != nil {
			fmt.Printf("Problem converting %s", chars[i])
		} else {
			digits = append(digits, converted)
		}
	}
	return digits
}
