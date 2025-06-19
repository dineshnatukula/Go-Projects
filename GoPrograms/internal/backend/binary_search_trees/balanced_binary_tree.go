package backend

import "golang.org/x/tour/tree"

func IsBalanced(root *tree.Tree) bool {
	_, balanced := checkHeight(root)
	return balanced
}

// Returns (height, isBalanced)
func checkHeight(node *tree.Tree) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkHeight(node.Left)
	rightHeight, rightBalanced := checkHeight(node.Right)

	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	return max(leftHeight, rightHeight) + 1, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
