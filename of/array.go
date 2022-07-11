// @Author shovenyuan
// @Date: 2022/7/1 16:34

// Package main ...
package main

func findRepeatNumber(nums []int) int {
	var m = map[int]int{} // k-num值，v-num值出现的次数
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 {
			return nums[i]
		} else {
			m[nums[i]] = 1
		}
	}
	return 0
}

func findRepeatNumber2(nums []int) int {
	i := 0
	for i < len(nums) {
		// 当nums[i] == i 时
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == i {
			return nums[i]
		}
		// 交换位置
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
