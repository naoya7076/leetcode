/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

// @lc code=start
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 1
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return depth
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return depth + max(leftDepth, rightDepth)
}

// @lc code=end
