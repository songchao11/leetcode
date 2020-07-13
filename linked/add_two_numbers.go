package linked

import (
	"fmt"
	"strconv"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

//示例：
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var numStr1 string
	var numStr2 string
	for {
		numStr1 = fmt.Sprintf("%d", l1.Val) + numStr1
		l1 = l1.Next
		if l1 == nil {
			break
		}
	}
	for {
		numStr2 = fmt.Sprintf("%d", l2.Val) + numStr2
		l2 = l2.Next
		if l2 == nil {
			break
		}
	}
	num1, err := strconv.Atoi(numStr1)
	if err != nil {
		return nil
	}
	num2, err := strconv.Atoi(numStr1)
	if err != nil {
		return nil
	}
	num := num1 + num2
	numStr := strconv.Itoa(num)
	nn := []rune(numStr)
	n, err := strconv.Atoi(string(nn[len(nn) - 1]))
	if err != nil {
		return nil
	}
	res := &ListNode{
		Val:  n,
		Next: nil,
	}
	if len(nn) > 1 {
		temp := res
		for i := len(nn) - 2;i >= 0;i-- {
			v, err := strconv.Atoi(string(nn[i]))
			if err != nil {
				return nil
			}
			temp.Next = &ListNode{
				Val:  v,
			}
			temp = temp.Next
		}
	}

	return res
}