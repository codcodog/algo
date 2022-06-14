package easy

import "testing"

type PalindromeData struct {
	Num      int
	Expected bool
}

func TestIsPalindrome(t *testing.T) {
	data := [3]PalindromeData{
		PalindromeData{
			Num:      121,
			Expected: true,
		},
		PalindromeData{
			Num:      -121,
			Expected: false,
		},
		PalindromeData{
			Num:      10,
			Expected: false,
		},
	}

	for _, item := range data {
		if IsPalindrome(item.Num) != item.Expected {
			t.Errorf("test is palindrome failed: %+v", item)
		}
		if IsPalindromeV1(item.Num) != item.Expected {
			t.Errorf("v1 test is palindrome failed: %+v", item)
		}
	}
}
