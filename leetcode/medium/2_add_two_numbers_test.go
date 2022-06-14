package medium

func ExampleAddTwoNumbers() {
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})
	l3 := AddTwoNumbers(l1, l2)
	l3.print()

	l1 = buildList([]int{0})
	l2 = buildList([]int{0})
	l3 = AddTwoNumbers(l1, l2)
	l3.print()

	l1 = buildList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = buildList([]int{9, 9, 9, 9})
	l3 = AddTwoNumbers(l1, l2)
	l3.print()

	// Output:
	// [7 0 8]
	// [0]
	// [8 9 9 9 0 0 0 1]
}

// build a list
func buildList(data []int) *ListNode {
	prehead := &ListNode{
		Val: -1,
	}

	prev := prehead
	for _, v := range data {
		node := &ListNode{
			Val: v,
		}
		prev.Next = node
		prev = prev.Next
	}

	return prehead.Next
}
