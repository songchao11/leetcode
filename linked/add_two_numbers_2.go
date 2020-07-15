package linked

//(中等)
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

//示例：
//输入：(2 -> 4 -> 3) +
//     (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

type ListNode struct {
	Val  int
	Next *ListNode
}

//误区：将链表转化成数字 然后相加 再转化成链表 (当链表表示的数字太大时，此时会溢出，则结果会错误)
//思路:本身链表表示的数字是倒序的 所以直接将链表的每位数字相加 过十进一 需要注意的是尾结点每次都要创建一个 会存在为空的现象
//所以每次先创建一个节点存数据 最后返回节点的Next节点
func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := new(ListNode)
	res := ln
	num := 0
	for l1 != nil || l2 != nil || num != 0 {
		ln.Next = new(ListNode)
		ln = ln.Next
		if l1 != nil {
			num = num + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num = num + l2.Val
			l2 = l2.Next
		}
		ln.Val = num % 10
		num = num / 10
	}
	return res.Next
}

//递归算法
func AddTwoNumbers2(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	nextNode := AddTwoNumbers2(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{
			Val:  sum,
			Next: nextNode,
		}
	} else {
		tempNode := &ListNode{
			Val:  1,
			Next: nil,
		}
		return &ListNode{
			Val:  sum - 10,
			Next: AddTwoNumbers2(nextNode, tempNode),
		}
	}

}
