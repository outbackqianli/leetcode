package main

import (
	"fmt"

	. "outback/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	removeNthFromEnd(head, 2)
}

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}
	dump := new(ListNode)
	dump.Next = head
	i, pre, cur := 0, dump, head
	for cur != nil {
		cur = cur.Next
		i++
		if i > n {
			pre = pre.Next
		}
	}
	fmt.Println(pre.Val)
	pre.Next = pre.Next.Next
	return dump.Next
}
