// @Author shovenyuan
// @Date: 2022/6/24 15:30

// Package main ...
package main

import "sort"

// 暴力法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 != 0 {
		// 奇数
		return float64(nums1[len(nums1)/2])
	} else {
		// 偶数
		return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
	}
}
