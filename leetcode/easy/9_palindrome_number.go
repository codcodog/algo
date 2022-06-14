package easy

import (
	"strconv"
)

// 转为字符串，检查是否回文
func IsPalindromeV1(x int) bool {
	if x < 0 {
		return false
	}

	strX := strconv.Itoa(x)
	s := []rune(strX)
	strLen := len(s)
	for index, value := range s {
		i := strLen - index - 1
		if index >= i {
			break
		}
		if value != s[i] {
			return false
		}
	}

	return true
}

// 将数字本身反转
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	dst, y := 0, x
	for y != 0 {
		dst = dst*10 + y%10
		y = y / 10
	}

	return dst == x
}
