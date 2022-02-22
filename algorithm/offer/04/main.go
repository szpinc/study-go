package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	nodes := []*ListNode{}

	next := head

	for next != nil {
		nodes = append(nodes, next)
		next = next.Next
	}

	values := []int{}

	for i := len(nodes) - 1; i >= 0; i-- {
		values = append(values, nodes[i].Val)
	}
	return values
}

type ListNode struct {
	Val  int
	Next *ListNode
}
