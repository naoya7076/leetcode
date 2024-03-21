/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // indexなので-1
	middle := 0
	for left <= right {
		middle = (right + left) / 2
		if target == nums[middle] {
			return middle
		} else if target > nums[middle] {
			left = middle + 1
		} else if target < nums[middle] {
			right = middle - 1
		}
	}
	return left
}

// @lc code=end
