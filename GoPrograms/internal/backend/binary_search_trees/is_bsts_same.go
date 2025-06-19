package backend

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func WalkTree(t1 *tree.Tree) []int {
	ch := make(chan int)
	var vals []int

	go func() {
		Walk(t1, ch)
		close(ch)
	}()

	for i := range ch {
		vals = append(vals, i)
	}

	return vals
}

// IsBSTsEqual determines whether the trees
// t1 and t2 contain the IsBSTsEqual values.
func IsBSTsEqual(t1, t2 *tree.Tree) bool {
	v1 := WalkTree(t1)
	v2 := WalkTree(t2)
	return reflect.DeepEqual(v1, v2)
}

// IsBTSsEqualWithStructures determines whether the trees
// t1 nad t2 contain the same structure along with the node values.
func IsBTSsEqualWithStructures(t1, t2 *tree.Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	return t1.Value == t2.Value &&
		IsBTSsEqualWithStructures(t1.Left, t2.Left) &&
		IsBTSsEqualWithStructures(t1.Right, t2.Right)
}

// Search(t1, val) returns true if the val exists in BST
// and false otherwise
func Search(t1 *tree.Tree, val int) bool {
	if t1 == nil {
		return false
	}

	if t1.Value == val {
		return true
	}

	if t1.Value < val {
		return Search(t1.Right, val)
	} else {
		return Search(t1.Left, val)
	}
}

// KthSmallestElement in BST return the Kth Smallest element if found
// and -1 otherwise

var cnt int

func KthSmallestElement(t1 *tree.Tree, k int) int {
	if t1 == nil {
		return -1
	}
	left := KthSmallestElement(t1.Left, k)
	if left != -1 {
		return left
	}
	cnt++
	if cnt == k {
		return t1.Value
	}
	return KthSmallestElement(t1.Right, k)
}

// KthSmallestElementFN in BST return the Kth Smallest element if found
// and -1 otherwise... With Helper Function

// var cnt int

func KthSmallestElementFN(t1 *tree.Tree, k int) int {
	var cnt int
	var inorder func(node *tree.Tree) int
	inorder = func(t1 *tree.Tree) int {
		if t1 == nil {
			return -1
		}
		left := inorder(t1.Left)
		if left != -1 {
			return left
		}
		cnt++
		if cnt == k {
			return t1.Value
		}
		return inorder(t1.Right)
	}
	kthSmallest := inorder(t1)
	return kthSmallest
}

// IsValidBST returns true if the BST is valid and false otherwise
func IsValidBST(t1 *tree.Tree) bool {
	return IsValidBSTFn(t1, nil, nil)
}

func IsValidBSTFn(t1 *tree.Tree, min *int, max *int) bool {
	if t1 == nil {
		return true
	}

	if min != nil && t1.Value <= *min {
		return false
	}
	if max != nil && t1.Value >= *max {
		return false
	}
	return IsValidBSTFn(t1.Left, min, &t1.Value) &&
		IsValidBSTFn(t1.Right, &t1.Value, max)
}

func insert(t *tree.Tree, Value int) *tree.Tree {
	if t == nil {
		return &tree.Tree{Value: Value}
	}

	if Value < t.Value {
		t.Left = insert(t.Left, Value)
	} else {
		t.Right = insert(t.Right, Value)
	}

	return t
}

func CreateBST() {
	t1 := &tree.Tree{Value: 0}
	t1 = insert(t1, 1)
	t1 = insert(t1, 3)
	t1 = insert(t1, -2)
	t1 = insert(t1, 2)
	t1 = insert(t1, 5)
	t1 = insert(t1, 6)
	t1 = insert(t1, 10)
	t1 = insert(t1, 7)
	t1 = insert(t1, 8)
	t1 = insert(t1, 9)

	fmt.Println(t1.String())

	t2 := &tree.Tree{Value: 0}
	t2 = insert(t2, -1)
	t2 = insert(t2, 6)
	t2 = insert(t2, 9)
	t2 = insert(t2, 2)
	t2 = insert(t2, 10)
	t2 = insert(t2, 1)
	t2 = insert(t2, 7)
	t2 = insert(t2, 3)
	t2 = insert(t2, 5)
	t2 = insert(t2, 8)

	fmt.Println(t2.String())

	fmt.Println("Traversing Tree t1", WalkTree(t1))
	// TraverseTree(t1)

	fmt.Println("Traversing Tree t2", WalkTree(t2))
	// TraverseTree(t2)

	fmt.Print("Trees t1 and t2 are ")
	if IsBSTsEqual(t1, t2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	fmt.Print("The BSTs have ")
	if IsBTSsEqualWithStructures(t1, t2) {
		fmt.Println("same structures")
	} else {
		fmt.Println("different structures")
	}

	if Search(t1, 10) {
		fmt.Println("Search successful, found 10")
	} else {
		fmt.Println("Search not successful, not found 10")
	}

	if Search(t1, -2) {
		fmt.Println("Search successful, found -2")
	} else {
		fmt.Println("Search not successful, not found -2")
	}

	for i := 1; i < 11; i++ {
		cnt = 0
		fmt.Printf("%dth SmallestElement from the BST t1 is %d\n", i, KthSmallestElement(t1, i))
	}
	for i := 1; i < 11; i++ {
		fmt.Printf("%dth SmallestElementFN from the BST t1 is %d\n", i, KthSmallestElementFN(t1, i))
	}

	if IsValidBST(t1) {
		fmt.Println("BST t1 is valid")
	} else {
		fmt.Println("BST t1 is not valid...")
	}

	if IsValidBST(t2) {
		fmt.Println("BST t2 is valid")
	} else {
		fmt.Println("BST t2 is not valid...")
	}
	fmt.Println("Done with Creation and Traversing the BST")
}
