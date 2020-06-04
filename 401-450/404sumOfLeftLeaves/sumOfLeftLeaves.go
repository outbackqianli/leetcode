package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
计算给定二叉树的所有左叶子之和。
示例：
    3
   / \
  9  20
    /  \
   15   7
在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
*/

// 注意,这里算的不是左节点,而不左叶子节点
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil  {
		return 0
	}

	ans := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if first.Left != nil && first.Left.Left == nil && first.Left.Right == nil {
			ans += first.Left.Val
		}
		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
	}
	return ans
}
