// @Author shovenyuan
// @Date: 2022/7/7 20:33

// Package main ...
package main

import (
	"strconv"
	"strings"
)

func combine(n int, k int) (ans [][]int) {
	var subSet []int
	var dfs = func(int) {}
	dfs = func(i int) {
		if len(subSet) == k {
			ans = append(ans, append([]int{}, subSet...))
			return
		}
		if i > n {
			return
		}
		subSet = append(subSet, i)
		dfs(i + 1)
		subSet = subSet[:len(subSet)-1]
		dfs(i + 1)
	}
	dfs(1)
	return ans
}

func restoreIpAddresses(s string) (ans []string) {
	var n = len(s)
	var subSet []string

	var dfs = func(int) {}
	dfs = func(curr int) {
		if curr == n && len(subSet) == 4 {
			ans = append(ans, strings.Join(subSet, "."))
			return
		}
		for i := curr + 1; i <= len(s); i++ {
			if ipNumCheck(s[curr:i]) {
				subSet = append(subSet, s[curr:i])
				dfs(i)
				subSet = subSet[:len(subSet)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func ipNumCheck(num string) bool {
	atoi, _ := strconv.Atoi(num)
	if atoi > 255 || (len(num) > 1 && num[0] == '0') {
		return false
	}
	return true
}
