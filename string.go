// @Author shovenyuan
// @Date: 2022/6/21 19:14

// Package main ...
package main

func singleNumber(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumberV2(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] ^= nums[i-1]
	}
	return nums[len(nums)-1]
}

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// 判断是否在栈顶
		if len(stack) != 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
