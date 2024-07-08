package main

import "strconv"

func ConverStringToList(str string) *ListNode {
	ans := ListNode{0, nil}
	iter := &ans
	for i := len(str) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(str[i]))
		if err == nil {
			iter.Next = &ListNode{num, nil}
		}
		iter = iter.Next
	}
	return ans.Next
}

func ConverListToString(lst *ListNode) string {
	str := ""
	iter := lst
	for iter != nil {
		str += strconv.Itoa(iter.Val)
		iter = iter.Next
	}
	return str
}

func Reverse(str1 string) string {
	j := len(str1) - 1
	ans := ""
	for i := j; i >= 0; i-- {
		ans += string(str1[i])
	}
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	str1 := ConverListToString(l1)
	str2 := ConverListToString(l2)
	str1 = Reverse(str1)
	str2 = Reverse(str2)
	num1, err1 := strconv.Atoi(str1)
	num2, err2 := strconv.Atoi(str2)
	if err1 == nil && err2 == nil {
		str3 := strconv.Itoa(num1 + num2)
		return ConverStringToList(str3)
	}
	return nil
}

