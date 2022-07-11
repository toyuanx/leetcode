// @Author shovenyuan
// @Date: 2022/7/8 17:35

// Package main ...
package main

import "strconv"

func evalRPN(tokens []string) int {
	// 保存数据与计算结果
	var stack = []int{}

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			tmp := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = tmp
		case "-":
			tmp := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = tmp
		case "*":
			tmp := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = tmp
		case "/":
			tmp := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = tmp
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	return stack[0]
}
