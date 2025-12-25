package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	sum1 := []string{}
	for cur1 != nil {
		sum1 = append(sum1, strconv.Itoa(cur1.Val))
		cur1 = cur1.Next
	}
	slices.Reverse(sum1)
	summ1 := strings.Join(sum1, "")
	summ11, err := strconv.Atoi(summ1)
	if err != nil {

	}

	cur2 := l2
	sum2 := []string{}
	for cur2 != nil {
		sum2 = append(sum2, strconv.Itoa(cur2.Val))
		cur2 = cur2.Next
	}
	slices.Reverse(sum2)
	summ2 := strings.Join(sum2, "")
	summ22, err := strconv.Atoi(summ2)
	if err != nil {

	}

	sum := summ11 + summ22
	r := strconv.Itoa(sum)
	re := strings.Split(r, "")
	res := sliceToList(re)
	fmt.Println(res)
	return res
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

func sliceToList(nums []string) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, v := range nums {
		int, err := strconv.Atoi(v)
		if err != nil {

		}
		curr.Next = &ListNode{Val: int}
		curr = curr.Next
	}
	return dummy.Next
}
