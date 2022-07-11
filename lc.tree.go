// @Author shovenyuan
// @Date: 2022/6/20 11:49
// 二叉树的一些题目

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

// 450. 删除二叉搜索树中的节点 https://leetcode.cn/problems/delete-node-in-a-bst/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 如果key大于val，则去右子树删除
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	// 如果key小于val，则去左子树删除
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}

	// 	如果key等于val，则需要判断左右子树情况
	if root.Val == key {
		// 左子树为空，直接返回右子树
		if root.Left == nil {
			return root.Right
		}
		// 右子树为空，直接返回左子树
		if root.Right == nil {
			return root.Left
		}

		// 左右子树都有值，将左子树移给最右子树左下角节点，右子树top节点成为根节点，从而实现节点的删除
		tmpNode := root.Right
		// 寻找最左节点
		for tmpNode.Left != nil {
			tmpNode = tmpNode.Left
		}
		tmpNode.Left = root.Left
		root = root.Right
	}

	return root
}

func isSymmetric(root *TreeNode) bool {
	// 左右同时遍历比较
	var helper func(*TreeNode, *TreeNode) bool
	helper = func(n1 *TreeNode, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		return helper(n1.Left, n2.Right) == helper(n1.Right, n2.Left)
	}
	return helper(root, root)
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	nodeList := strings.Split(data, ",")
	index := [1]int{0} // 数组指针传递

	var dfs = func(nodeList []string, index [1]int) *TreeNode { return nil }
	dfs = func(nodeList []string, index [1]int) *TreeNode {
		nodeValStr := nodeList[index[0]]
		index[0]++
		if nodeValStr == "null" {
			return nil
		}
		nodeVal, _ := strconv.Atoi(nodeValStr)
		node := &TreeNode{Val: nodeVal}
		node.Left = dfs(nodeList, index)
		node.Right = dfs(nodeList, index)
		return node
	}
	return dfs(nodeList, index)
}
