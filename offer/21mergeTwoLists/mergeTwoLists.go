package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dump := new(ListNode)
	if l1.Val < l2.Val {
		dump.Next = l1
		dump.Next.Next = mergeTwoLists(l1.Next, l2)
		return dump.Next
	} else {
		dump.Next = l2
		dump.Next.Next = mergeTwoLists(l1, l2.Next)
		return dump.Next
	}
}
