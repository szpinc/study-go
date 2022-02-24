package main

import "fmt"

func main() {
	n3 := ListNode{
		Val: 3,
	}
	n2 := ListNode{
		Val:  2,
		Next: &n3,
	}
	n1 := ListNode{
		Val:  1,
		Next: &n2,
	}
	node := ReverseList(&n1)

	next := node

	values := []int{}
	for next != nil {
		values = append(values, next.Val)
		next = next.Next
	}

	fmt.Printf("%v", values)
}

// 反转链表
func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}

	var newHead *ListNode = nil

	for pHead != nil {
		// 保留下一个节点
		next := pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = next
	}
	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
