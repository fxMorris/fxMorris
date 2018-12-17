package leetcode

//ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	p := &ListNode{}
	re := p
	for {
		bit1, bit2 := 0, 0
		if l1 != nil {
			bit1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			bit2 = l2.Val
			l2 = l2.Next
		}
		bit := bit1 + bit2 + flag
		flag = 0
		if bit >= 10 {
			bit -= 10
			flag = 1
		}
		p.Val = bit
		//链表已到尾，且没有进位，输出结果
		if flag == 0 && l1 == nil && l2 == nil {
			return re
		}
		p.Next = &ListNode{}
		p = p.Next
	}
}
