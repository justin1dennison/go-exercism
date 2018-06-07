package bob

import "strings"

func Hey(remark string) string {
	if isSilence(remark) {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) && isYelling(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isYelling(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion(remark) {
		return "Sure."
	}
	return "Whatever."
}

func isSilence(remark string) bool {
	chars := []string{" ", "\t", "\n", "\r"}
	var result string = remark
	for _, char := range chars {
		result = strings.Replace(result, char, "", -1)
	}
	return result == ""

}

func isQuestion(remark string) bool {
	trimmed := strings.Trim(remark, " ")
	lastChar := string(trimmed[len(trimmed)-1])
	return lastChar == "?"
}

func isYelling(remark string) bool {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return strings.ToUpper(remark) == remark && strings.ContainsAny(remark, letters)
}
