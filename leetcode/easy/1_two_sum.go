package easy

func TwoSum(nums []int, target int) (indexes []int) {
	if len(nums) < 2 {
		return nil
	}

	cache := make(map[int]int)
	for index, value := range nums {
		if v, ok := cache[value]; ok {
			return []int{v, index}
		}
		cache[target-value] = index
	}

	return nil
}
