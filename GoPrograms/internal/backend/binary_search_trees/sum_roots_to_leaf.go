package backend

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func AppendRootsToLeaf(t *tree.Tree) int {
	sum := 0
	var dfs func(t *tree.Tree, sum int) int
	dfs = func(t *tree.Tree, sum int) int {
		if t == nil {
			return 0
		}
		sum = sum*10 + t.Value
		if t.Left == nil && t.Right == nil {
			return sum
		}
		return dfs(t.Left, sum) + dfs(t.Right, sum)
	}
	return dfs(t, sum)
}

func CreateBSTForTesing() *tree.Tree {
	t1 := &tree.Tree{Value: 4}
	t1 = insert(t1, 9)
	// t1 = insert(t1, 0)
	t1 = insert(t1, 5)
	// t1 = insert(t1, 1)
	// t1 = insert(t1, -2)
	// t1 = insert(t1, 2)
	// t1 = insert(t1, 5)
	// t1 = insert(t1, 6)
	// t1 = insert(t1, 10)
	// t1 = insert(t1, 7)
	// t1 = insert(t1, 8)
	// t1 = insert(t1, 9)
	fmt.Println(t1.String())
	return t1
}

func MinDepth(root *tree.Tree) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}
	return min(MinDepth(root.Left), MinDepth(root.Right)) + 1
}
