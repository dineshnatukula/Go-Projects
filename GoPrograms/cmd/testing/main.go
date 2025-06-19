package main

import (
	be "myproject/internal/backend/maps"
)

func main() {

	// workerpools.NewDeepPool(10, 10)
	// workerpools.WorkerPoolDemo()
	// workerpools.WorkerPoolWithContext()

	// Demostrating Database Deadlock
	// backend.DatabaseDeadLockDemo()
	// Demostrating Database Deadlock along with retry mechanism
	// backend.DatabaseDeadLockRetryMech()

	// gracefulshutdown.GracefulShutdownDemo()
	// wp.WPDemo()

	// context.ContextDoneMultiple()

	// Interface demo
	// backend.InterfaceDemo()
	// backend.TypeAssertion()
	// backend.PrintAnything("hello")

	// ReflectDemo Demo
	// reflect.ReflectDemo()

	// Demos Map
	// backend.MyMap()

	// Explaining Go Routines with Wait Groups
	// backend.GoRoutine()

	// Demonstrating Context:
	// context.ContextSuccess()

	// Explaining Go Routines with Select Case
	// backend.GoRoutineWithSelect()

	// Demonstrating Sync Mutex
	// backend.SyncMutex()

	// Type Parameters
	// backend.TypeParams()

	// Creating BST
	// backend.CreateBST()

	// Create Single Linked List
	// fmt.Println("Create Single Linked List")
	// be.CreateList()

	// Sum root to leaf for BST
	// fmt.Println()
	// fmt.Println("Sum Root To Leaf in BST", backend.AppendRootsToLeaf(backend.CreateBSTForTesing()))

	// get single digit from n
	// fmt.Println("Get Single Digit Sum", strings.GetSingleDigit(28))

	// get excel column title based on val
	// fmt.Println("Get Excel Column Title", strings.ExcelColumnTitle(1))
	// fmt.Println("Get Excel Column Title", strings.ExcelColumnTitle(26))
	// fmt.Println("Get Excel Column Title", strings.ExcelColumnTitle(27))
	// fmt.Println("Get Excel Column Title", strings.ExcelColumnTitle(28))
	// fmt.Println("Get Excel Column Title", strings.ExcelColumnTitle(56))
	// fmt.Println("Get Excel Column Title", strings.ExcelColumnTitle(701))

	// Get the minimum depth of BST
	// fmt.Println("Min Depth in BST", backend.MinDepth(backend.CreateBSTForTesing()))

	// Check if the Binary tree is balanced?
	// fmt.Println("Is Balacend Binary Tree", backend.IsBalanced(backend.CreateBSTForTesing()))

	// Read data from csv
	// fmt.Println("Read data from csv ", strings.ReadTests())

	// Sync Map Demo.
	be.SyncMapDemo()

}
