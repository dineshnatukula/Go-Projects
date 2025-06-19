package singlelinkedlist

import (
	"fmt"
	"myproject/internal/models"
)

func CreateNode(value int) *models.List[int] {
	tmp := &models.List[int]{Value: value, Next: nil}
	if tmp == nil {
		panic("Error in creating the node")
	}
	return tmp
}

func InsertFirst(first, last *models.List[int], value int) {
	if first == nil {
		first = CreateNode(value)
		last = first
	} else {
		tmp := CreateNode(value)
		tmp.Next = first
		first = tmp
	}
}

func InsertLast(first, last *models.List[int], value int) {
	if first == nil {
		InsertFirst(first, last, value)
	} else {
		tmp := CreateNode(value)
		last.Next = tmp
		last = tmp
	}
}

func CreateList() {
	var first *models.List[int]
	var last *models.List[int]
	InsertFirst(first, last, 1)
	InsertLast(first, last, 2)
	InsertFirst(first, last, 3)
	InsertLast(first, last, 4)
	InsertFirst(first, last, 5)
	InsertFirst(first, last, 6)
	InsertLast(first, last, 7)
	TraverseList(first)
}

func HasCycle(first *models.List[int]) bool {
	slow := first
	fast := first

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func TraverseList(first *models.List[int]) {
	for tmp := first; tmp != nil; tmp = tmp.Next {
		fmt.Printf("%d -> ", tmp.Value)
	}
}
