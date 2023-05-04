package main

// https://leetcode.com/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	previousNode := head
	if head != nil {
		currentNode := head.Next
		for currentNode != nil {
			if currentNode.Val == val {
				previousNode.Next = currentNode.Next
			} else {
				previousNode = currentNode
			}
			currentNode = currentNode.Next

		}
		if head.Val == val {
			head = head.Next
		}
	}
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return dummy.Next
}
