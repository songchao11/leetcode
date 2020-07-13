package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int64
	Next *ListNode
}
func main()  {

	//nodeA3 := &ListNode{
	//	Val:  3,
	//	Next: nil,
	//}
	//nodeA2 := &ListNode{
	//	Val:  4,
	//	Next: nodeA3,
	//}
	//l1 := &ListNode{
	//	Val:  2,
	//	Next: nodeA2,
	//}
	//nodeB3 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//nodeB2 := &ListNode{
	//	Val:  6,
	//	Next: nodeB3,
	//}
	//l2 := &ListNode{
	//	Val:  5,
	//	Next: nodeB2,
	//}

	//[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	//[5,6,4]

	arr1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	arr2 := []int{5, 6, 4}
	l1 := &ListNode{Val: int64(arr1[0])}
	temp1 := l1
	for i := 1;i < len(arr1);i++{
		temp1.Next = &ListNode{Val:int64(arr1[i])}
		temp1 = temp1.Next
	}
	l2 := &ListNode{Val: int64(arr2[0])}
	temp2 := l2
	for i := 1;i < len(arr2);i++{
		temp2.Next = &ListNode{Val:int64(arr2[i])}
		temp2 = temp2.Next
	}

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
		//return nil
	}
	fmt.Println(num1)
	num2, err := strconv.Atoi(numStr2)
	if err != nil {
		//return nil
	}
	num := num1 + num2
	numStr := strconv.Itoa(num)
	nn := []rune(numStr)
	n, err := strconv.Atoi(string(nn[len(nn) - 1]))
	if err != nil {
		//return nil
	}
	res := &ListNode{
		Val:  int64(n),
		Next: nil,
	}
	if len(nn) > 1 {
		temp := res
		for i := len(nn) - 2;i >= 0;i-- {
			v, err := strconv.Atoi(string(nn[i]))
			if err != nil {
				//return nil
			}
			temp.Next = &ListNode{
				Val:  int64(v),
			}
			temp = temp.Next
		}
	}

	//return res
	for {
		fmt.Println(res.Val)
		res = res.Next
		if res == nil {
			break
		}
	}
}
