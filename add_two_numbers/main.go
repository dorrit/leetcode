package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	l1 := n1
	l2 := n1

	AddTwoNumbers(l1, l2)
}
