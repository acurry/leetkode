package main

const NULL_NODE = -9999

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return false
}

func buildTree(nodes []int) *TreeNode {
	var root TreeNode
	for _, n := range nodes {
		if n != NULL_NODE {
			root.Val = n
			recur(root.Left)
			recur(root.Right)
		}
	}

	return &root
}

func recur(n *TreeNode)
