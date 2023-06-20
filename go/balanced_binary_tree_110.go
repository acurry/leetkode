package main

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	isHeightWithinOne := lh-rh == 0 || lh-rh == -1 || lh-rh == 1

	if isHeightWithinOne && IsBalanced(root.Left) && IsBalanced(root.Right) {
		return true
	}

	return false
}
