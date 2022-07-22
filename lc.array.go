// @Author shovenyuan
// @Date: 2022/6/24 15:30

// Package main ...
package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(Parition([]int{3, 2, 1, 5, 6, 4}, 0, 5))
	fmt.Println(threeSum([]int{-2, -3, 0, 0, -2}))
}

func longestConsecutive(nums []int) int {
	var maxLen int
	// 先定义hash表，key为nums中的数字，顺带去重
	var hash = make(map[int]bool)
	for _, v := range nums {
		hash[v] = true
	}
	fmt.Println(hash[100], hash[7])
	// 循环hash表，找到起始位置x: x的判断条件是判断x-1存不存在
	for k := range hash {
		if _, ok := hash[k-1]; !ok {
			end := k
			for hash[end+1] {
				end++
			}
			if end-k+1 > maxLen {
				maxLen = end - k + 1
			}
		}
	}
	return maxLen
}

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

func Parition(nums []int, l, r int) int {
	if l >= r {
		return -1
	}
	pivot := nums[l]
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] < pivot {
			l++
		}
		nums[r] = nums[l]
	}
	// 循环结束，l与r相等
	// 确定基准元素pivot在数组中的位置
	nums[l] = pivot
	return l
}

func threeSum(nums []int) (ans [][]int) {
	// 先排序，然后双指针
	sort.Ints(nums)
	// fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		// 数组已排序，所以当nums[i]大于0时，直接返回
		if nums[i] > 0 {
			break
		}
		// fmt.Println(i)
		// 左右指针索引，初始值
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			// 和大于0， r值太大，-1且跳过所有重复的num
			if sum > 0 {
				r--
				for r > 0 && nums[r+1] == nums[r] {
					r--
				}
				continue
			}

			// 和大于0， l值太小，+1且跳过所有重复的num
			if sum < 0 {
				l++
				for l < len(nums)-1 && nums[l] == nums[l-1] {
					l++
				}
				continue
			}

			// 和为0， 去除重复解
			ans = append(ans, []int{nums[i], nums[l], nums[r]})
			for l < r && nums[l] == nums[l+1] {
				l++
			}
			for l < r && nums[r-1] == nums[r] {
				r--
			}
			l++
			r--
		}
		// i值重复
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func findUnsortedSubarray(nums []int) int {

	return 0
}

func rotate(matrix [][]int) {
	// 辅助矩阵
	var tmp [][]int
	var n = len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[i][j] = matrix[j][n-i-1]
		}
	}
	copy(matrix, tmp)
}
