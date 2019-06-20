package main

import "fmt"

func main() {
	input := "23"
	fmt.Println(letterCombinations(input)) // expected: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
}

var digitsMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	result := []string{""}
	for i := 0; i < len(digits); i++ {
		chars := digitsMap[string(digits[i])]
		tmp := make([]string, 0)
		for _, char := range chars {
			for _, old := range result {
				tmp = append(tmp, fmt.Sprintf("%s%s", old, char))
			}
		}
		result = tmp
	}
	return result
}
