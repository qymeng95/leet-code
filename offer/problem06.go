package offer

import . "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)

		head = head.Next
	}

	for i, j := 0, len(res); i < j; i++ {
		j = len(res) - (i + 1)
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func ReversePrint2(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append([]int{head.Val}, res...)

		head = head.Next
	}

	return res
}

func TestReversePrint() {
	n := &ListNode{1, &ListNode{2, nil}}
	res := ReversePrint(n)
	res2 := ReversePrint2(n)

	Println(res)
	Println(res2)
}
