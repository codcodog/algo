package easy

// f(n) = f(n-1) + f(n-2)
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var temp int
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		temp = a + b
		a, b = b, temp
	}

	return temp
}

var cache = make(map[int]int)

// 使用递归 f(n) = f(n-1) + f(n-2)
func ClimbStairsV1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	steps, ok := cache[n]
	if ok {
		return steps
	}

	steps = ClimbStairsV1(n-1) + ClimbStairsV1(n-2)
	cache[n] = steps

	return steps
}
