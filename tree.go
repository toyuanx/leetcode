// @Author shovenyuan
// @Date: 2022/6/20 11:49

// Package main ...
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) (ans []string) {
	var binaryTree func(*TreeNode)
	binaryTree = func(node *TreeNode) {
		fmt.Println(node)
		if node == nil {
			return
		}
		binaryTree(node.Left)
		binaryTree(node.Right)
		// 叶子节点
		if node.Left == nil && node.Right == nil {
			ans = append(ans, strconv.Itoa(node.Val))
			return
		}
		// 非叶子节点
		ans[len(ans)-1] = strconv.Itoa(node.Val) + "->" + ans[len(ans)-1]
	}
	binaryTree(root)
	return ans
}

func rightSideView(root *TreeNode) (ans []int) {
	var loop func(*TreeNode, *[]int, int)
	loop = func(node *TreeNode, ans *[]int, depth int) {
		if node == nil {
			return
		}
		if depth == len(*ans) {
			*ans = append(*ans, node.Val)
		}
		loop(node.Right, ans, depth+1)
		loop(node.Left, ans, depth+1)
	}

	loop(root, &ans, 0)
	return
}

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	var level func(*TreeNode, *[][]int, int)
	level = func(node *TreeNode, ans *[][]int, l int) {
		if node == nil {
			return
		}
		if len(*ans) > l {
			(*ans)[l] = append((*ans)[l], node.Val)
		} else {
			*ans = append(*ans, []int{node.Val})
		}

		level(node.Left, ans, l+1)
		level(node.Right, ans, l+1)
		return
	}

	level(root, &ans, 0)
	return
}

func maxDepth(root *TreeNode) int {
	q := []*TreeNode{root}
	l := 0
	for len(q) != 0 {
		node := q[0]
		q := q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		l++
	}
	return l
}

func levelOrder2(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		ql := len(q)
		subAns := []int{}
		for ql > 0 {
			node := q[0]
			q = q[1:]
			subAns = append(subAns, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			ql--
		}
		ans = append(ans, subAns)
	}
	return ans
}

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ans []string
	var dfs = func(*TreeNode) {}
	dfs = func(node *TreeNode) {
		if node == nil {
			ans = append(ans, "null")
			return
		}
		ans = append(ans, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	list := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		curr := list[0]
		list = list[1:]
		if curr == "null" {
			return nil
		}
		val, _ := strconv.Atoi(list[0])
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}
