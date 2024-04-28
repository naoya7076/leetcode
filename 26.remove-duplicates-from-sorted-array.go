/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastUniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastUniqueIndex] {
			lastUniqueIndex++
			nums[lastUniqueIndex] = nums[i]
		}
	}
	return lastUniqueIndex + 1
}

// @lc code=end
