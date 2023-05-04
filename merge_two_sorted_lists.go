package main

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res, cur1, cur2, resF *ListNode
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		resF = list1
		cur2 = list2
		cur1 = list1.Next
	} else {
		resF = list2
		cur2 = list2.Next
		cur1 = list1
	}

	res = resF

	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			res.Next = cur2
			break
		}

		if cur2 == nil {
			res.Next = cur1
			break
		}

		if cur1.Val < cur2.Val {
			res.Next = cur1
			cur1 = cur1.Next
		} else {
			res.Next = cur2
			cur2 = cur2.Next
		}
		res = res.Next
	}
	return resF
}
