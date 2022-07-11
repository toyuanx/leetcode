// @Author shovenyuan
// @Date: 2022/6/21 19:14

// Package main ...
package main

import "fmt"

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

func lengthOfLongestSubstring(s string) int {
	start, tmp, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		tmp = i - start + 1
		for j := start; j < i; j++ {
			fmt.Println(i, j, start, string(s[i]), string(s[j]))
			if s[i] == s[j] {
				tmp = i - start
				start = j + 1
				break
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func countSubstrings(s string) int {
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		// j = 0 为奇数回文串，j = 1为偶数回文串
		for j := 0; j <= 1; j++ {
			l := i
			r := i + j
			for l > 0 && r < n && s[l] == s[r] {
				count++
				l--
				r++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
