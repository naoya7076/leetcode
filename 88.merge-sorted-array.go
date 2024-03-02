/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
package main

import (
	"sort"
)

// @lc code=start

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, num := range nums2 {
		nums1[m+i] = num
	}
	sort.Ints(nums1)
}

// @lc code=end
