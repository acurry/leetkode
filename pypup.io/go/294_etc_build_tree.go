package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 294
func build_tree_i() *TreeNode {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	return &root
}

// 295
func build_tree_ii() *TreeNode {
	root := TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 8,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	return &root
}

// 296
func build_tree_iii() *TreeNode {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	return &root
}

// 299
func right_left_root(root *TreeNode) int {
	if root == nil {
		return -1
	}

	l := root.Left
	if l == nil {
		return -1
	}

	r := l.Right
	if r == nil {
		return -1
	}

	return r.Val
}

// 300
func pre_order(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	pre_order(root.Left)
	pre_order(root.Right)
}

// 301
func in_order(root *TreeNode) {
	if root == nil {
		return
	}
	in_order(root.Left)
	fmt.Println(root.Val)
	in_order(root.Right)
}

// 302
func post_order(root *TreeNode) {
	if root == nil {
		return
	}
	post_order(root.Left)
	post_order(root.Right)
	fmt.Println(root.Val)
}

// 303
func count_tree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + count_tree(root.Left) + count_tree(root.Right)
}

// 304
func sum_tree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + sum_tree(root.Left) + sum_tree(root.Right)
}

// 305
func sum_tree_node(root *TreeNode) int {
	return root.Val + root.Left.Val + root.Right.Val
}

// 307
func max_tree(root *TreeNode) int {
	if root == nil {
		return -1
	} else if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return max(root.Val, max_tree(root.Left), max_tree(root.Right))
}

// 308
/* Found a bug in this one on the site */
/*
Test 5
Input:
["20(21)(23(24)(25(29)))"]
Output:
-1
Expected:
20
*/
func min_tree(root *TreeNode) int {
	if root == nil {
		return -1
	} else if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return min(root.Val, max_tree(root.Left), max_tree(root.Right))
}

// 309
func exists_tree(root *TreeNode, x int) bool {
	if root == nil {
		return false
	}

	return root.Val == x || exists_tree(root.Left, x) || exists_tree(root.Right, x)
}

// 310
func even_count_tree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	is_even := 0
	if root.Val%2 == 0 {
		is_even = 1
	}

	return is_even + even_count_tree(root.Left) + even_count_tree(root.Right)
}

func min(f ...int) int {
	i := make([]int, 3)
	copy(i, f)
	sort.Ints(i)
	return i[0]
}

func max(a, b, c int) int {
	i := []int{a, b, c}
	sort.Ints(i)
	return i[len(i)-1]
}

// https://leetcode.com/problems/invert-binary-tree/submissions/984726661/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		return &TreeNode{
			Val:   root.Val,
			Left:  invertTree(root.Right),
			Right: invertTree(root.Left),
		}
	}
}
