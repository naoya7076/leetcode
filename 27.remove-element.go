/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
package main

// @lc code=start
func removeElement(nums []int, val int) int {
	k := 0
	for i, num := range nums {
		if num != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// @lc code=end
