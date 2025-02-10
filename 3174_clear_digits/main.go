package main

import (
	"fmt"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func clearDigits(s string) string {
	var result []rune

	for _, c := range s {
		if isDigit(c) {
			if len(result) > 0 && !isDigit(result[len(result)-1]) {
				result = result[:len(result)-1]
			} else {
				result = append(result, c)
            }

		} else {
			result = append(result, c)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(clearDigits("abc"))
	fmt.Println(clearDigits("cb34"))
	fmt.Println(clearDigits("12cb34"))
}
