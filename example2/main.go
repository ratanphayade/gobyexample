package main

import (
	"fmt"
)

var digitToLetters = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations := []string{""}

	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		letters := digitToLetters[digit]
		newCombinations := []string{}

		for _, letter := range letters {
			for _, combo := range combinations {
				newCombinations = append(newCombinations, combo+letter)
			}
		}

		combinations = newCombinations
	}

	return combinations
}

func main() {
	digits := "23"
	result := letterCombinations(digits)
	fmt.Println(result)
}
