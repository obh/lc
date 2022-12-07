package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var selNode = &ListNode{}
	var selIndex = -1
	merged := &ListNode{}
	origMerged := merged
	for {
		for i, currNode := range lists {
			if currNode != nil && (selIndex == -1 || selNode.Val > currNode.Val) {
				selIndex = i
				selNode = currNode
			}
		}
		if selIndex == -1 {
			break
		}
		merged.Next = &ListNode{selNode.Val, nil}
		merged = merged.Next
		lists[selIndex] = selNode.Next
		selIndex = -1
	}
	return origMerged.Next
}

func main() {
	l1_3 := &ListNode{5, nil}
	l1_2 := &ListNode{4, l1_3}
	l1_1 := &ListNode{1, l1_2}

	l2_3 := &ListNode{4, nil}
	l2_2 := &ListNode{3, l2_3}
	l2_1 := &ListNode{1, l2_2}

	l3_2 := &ListNode{0, nil}
	l3_1 := &ListNode{0, l3_2}

	lists := []*ListNode{l1_1, l2_1, l3_1}
	// empty := &ListNode{}
	// lists := []*ListNode{}
	merged := mergeKLists(lists)

	// for e := merged; e != nil; e = e.Next {
	// 	fmt.Print(e.Val)
	// 	fmt.Print(" -> ")
	// }
	fmt.Println(merged)
}
