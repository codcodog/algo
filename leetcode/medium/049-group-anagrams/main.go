package main

import "fmt"
import "sort"
import "strings"

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	result := make([][]string, 0)
	cache := make(map[string][]string)
	for _, value := range strs {
		sortStr := sortString(value)
		cache[sortStr] = append(cache[sortStr], value)
	}

	for _, lst := range cache {
		result = append(result, lst)
	}
	return result
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
