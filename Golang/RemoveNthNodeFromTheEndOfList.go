package main

type ListNode struct {
	Val int
	Next *ListNode
}

func lenghtOfList(head *ListNode) int {
	iter := head
	size := 0
	for iter != nil {
        iter = iter.Next
        size++
    }
	return size
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    size := lenghtOfList(head) - n
	if head == nil || size < 0 {
        return nil
    }
	if size == 0 {
		head = head.Next
		return head
	}
	iter := head
	var prev *ListNode
	count := 0
	for iter != nil {
		if (count == size) {
			prev.Next = iter.Next
			break
		}
		prev = iter
		iter = iter.Next
		count++
	}
	return head
}