package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	reorderList(root)
	PrintListNode(root)
}

/*

给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/

func reorderList(head *ListNode) {
	//fmt.Println("root value ", head.Val)
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	tailPre := head
	// 这里可以使用map优化
	for tailPre.Next.Next != nil {
		tailPre = tailPre.Next
	}

	tailPre.Next.Next = head.Next
	head.Next = tailPre.Next
	tailPre.Next=nil
	reorderList(head.Next.Next)
}
