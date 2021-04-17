package codecn

import "fmt"

//You are given two non-empty linked lists representing two non-negative integer
//s. The digits are stored in reverse order, and each of their nodes contains a si
//ngle digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the nu
//mber 0 itself.
//
//
// Example 1:
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//
//
// Example 2:
//
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//
//
// Example 3:
//
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//
// Constraints:
//
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading
// zeros.
//
// Related Topics 递归 链表 数学
// 👍 6003 👎 0

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lhead := &ListNode{
		Val:  0,
		Next: nil,
	}
	lpos := &ListNode{
		Val:  0,
		Next: nil,
	}
	lpriv := &ListNode{
		Val:  0,
		Next: nil,
	}

	var index, v1, v2 int = 0, 0, 0
	var sum, carry, remainder int
	for ; l1 != nil || l2 != nil; index++ {
		if nil == l1 {
			v1 = 0
		} else {
			v1 = l1.Val
		}

		if nil == l2 {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		sum = v1 + v2 + carry
		remainder = sum % 10

		if index == 0 {
			lhead.Val = remainder
			lhead.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			lpriv = lhead
			lpos = lhead.Next
		} else {
			lpos.Val = remainder
			lpos.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			lpriv = lpos
			lpos = lpos.Next
		}

		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		lpos.Val = carry
	} else {
		lpriv.Next = nil
	}

	return lhead
}

func TestP2() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  9,
		Next: nil,
	}

	lres := addTwoNumbers(l1, l2)
	for lres != nil {
		fmt.Println(lres.Val)
		lres = lres.Next
	}
}

//leetcode submit region end(Prohibit modification and deletion)
