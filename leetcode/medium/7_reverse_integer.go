package medium

import (
	"math"
)

func ReverseInt(n int) int {
	var target int
	for n != 0 {
		m := n % 10
		n = n / 10
		target = target*10 + m

		if target > math.MaxInt32 || target < math.MinInt32 {
			return 0
		}
	}

	return target
}
