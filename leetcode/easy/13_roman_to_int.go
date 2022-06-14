package easy

import (
	"strings"
)

var mapTable = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// 利用字符匹配
func RomanToIntV1(s string) int {
	var total int
	var lastC string
	data := strings.Split(s, "")
	for i := 0; i < len(data); i++ {
		c := data[i]
		switch c {
		case "I":
			total += mapTable[c]
		case "V", "X":
			if lastC == "I" {
				total = total + mapTable[c] - 2*mapTable[lastC]
			} else {
				total = total + mapTable[c]
			}
		case "L", "C":
			if lastC == "X" {
				total = total + mapTable[c] - 2*mapTable[lastC]
			} else {
				total = total + mapTable[c]
			}
		case "D", "M":
			if lastC == "C" {
				total = total + mapTable[c] - 2*mapTable[lastC]
			} else {
				total = total + mapTable[c]
			}
		}
		lastC = c
	}

	return total
}

var symbolValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	total, n := 0, len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}

	return total
}
