package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) findMiddleNode() *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var fast, slow = head, head
	// когда fast достигнет конца списка
	// slow будет указывать на середину списка
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	a := ListNode{
		Val:  5,
		Next: nil,
	}

	b := ListNode{
		Val:  21,
		Next: &a,
	}

	c := ListNode{
		Val:  56,
		Next: &b,
	}

	d := ListNode{
		Val:  2334,
		Next: &c,
	}

	node := d.findMiddleNode()

	fmt.Println(node.Val)
}
