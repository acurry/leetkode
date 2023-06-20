package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SortedArrayToBST(nums []int) *TreeNode {
	return recur(nums, 0, len(nums)-1)
}

func recur(nums []int, start, finish int) *TreeNode {
	if start > finish {
		return nil
	}

	midpoint := start + (finish-start)/2
	root := &TreeNode{Val: nums[midpoint]}

	root.Left = recur(nums, start, midpoint-1)
	root.Right = recur(nums, midpoint+1, finish)

	return root
}
