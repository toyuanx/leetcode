// @Author shovenyuan
// @Date: 2022/6/23 14:50

// Package main ...
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// 链接：https://leetcode.cn/leetbook/read/top-interview-questions/xm77tm/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func searchMatrix(matrix [][]int, target int) bool {
	// 行宽
	rows := len(matrix)
	// 列宽
	cols := len(matrix[0])

	// 左下角开始
	i := 0
	j := cols - 1
	for i < rows && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func isPalindrome(s string) bool {
	// 筛选字符串，组成新字符串
	ns := ""
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			ns += string(v)
		}
	}
	ns = strings.ToLower(ns)
	for i := 0; i < len(ns)/2; i++ {
		if ns[i] != ns[len(ns)-1-i] {
			return false
		}
	}
	return true
}

func sortArray(nums []int) []int {
	// 冒泡排序，比较交换，稳定算法，时间O(n^2), 空间O(1)
	// 每一轮遍历，将该轮最大值放到后面，同时小的往前冒
	// 从而形成后部是有序区
	// compare and swap
	for i := 0; i < len(nums); i++ {
		// 适当剪枝，len()-i到最后的部分都是有序区，避免再排
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}

func sortArrayQuickSort(nums []int) []int {
	// 快速排序
	if len(nums) <= 1 {
		return nums
	}
	// 基准数
	baseNum := nums[0]
	// 左、右边数组，其中左边为小于基准数的数组，右边为大于基准书的数组
	var left, right []int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= baseNum {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	left = sortArray(left)
	right = sortArray(right)

	// 返回左数组+baseNum+右数组
	return append(append(left, baseNum), right...)
}
