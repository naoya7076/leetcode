/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt64
	prev := -1

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)

		if prev >= 0 {
			diff := node.Val - prev
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node.Val

		inOrder(node.Right)
	}
	inOrder(root)
	return minDiff
}

// @lc code=end
