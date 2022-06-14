package easy

func ExampleMergeTwoLists() {
	l1 := buildList([]int{1, 2, 4})
	l2 := buildList([]int{1, 3, 4})

	l3 := MergeTwoLists(l1, l2)
	l3.print()
	// Output:
	// [1 1 2 3 4 4]
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
