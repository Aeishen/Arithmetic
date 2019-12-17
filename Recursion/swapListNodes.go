/*
   @File : swapListNodes
   @Author: Aeishen
   @Date: 2019/12/16 22:00
   @Description:递归实现两两交换链表中的节点
*/

/*
题目：
	给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例：
	给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

	head := swapPairs(&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,nil}}}})
	fmt.Println(head)

	head = swapPairs(&ListNode{1,nil})
	fmt.Println(head)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil{
		return head
	}

	temp := head.Next
	head.Next = swapPairs(temp.Next)
	temp.Next = head
	return temp
}