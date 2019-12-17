/*
   @File: ReverseList
   @Author: Aeishen
   @Date: 2019/12/17 10:23
   @Description: 反转链表
*/

/*
题目：
	反转一个单链表。

示例:
	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL

进阶:
	你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

思路：
	1.迭代：
	在遍历列表时，将当前节点的 next 指针改为指向前一个元素。由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。在更改
    引用之前，还需要另一个指针来存储下一个节点。不要忘记在最后返回新的头引用！
	复杂度分析
		时间复杂度：O(n)O(n) 。 假设 nn 是列表的长度，时间复杂度是 O(n)O(n)。
		空间复杂度：O(1)O(1) 。

	2.递归：
    递归版本稍微复杂一些，其关键在于反向工作。假设列表的其余部分已经被反转，现在我该如何反转它前面的部分？假设列表为：n1 → … → nk-1 → nk → nk+1 → … → nm → nil
	若从节点 nk+1 到 nm 已经被反转，而我们正处于 nk。n1 → … → nk-1 → nk → nk+1 ← … ← nm。我们希望 nk+1 的下一个节点指向 nk。所以，nk.next.next = nk;
	要小心的是 n1 的下一个必须指向 nil 。如果你忽略了这一点，你的链表中可能会产生循环。如果使用大小为 2 的链表测试代码，则可能会捕获此错误。
*/

package main

type ListNode1 struct {
     Val int
     Next *ListNode1
}

func main() {
	newList := &ListNode1{1,&ListNode1{2,&ListNode1{3,&ListNode1{4,&ListNode1{5,nil}}}}}
	reverseList1(newList)
}

// 迭代1(瞎写)
func reverseList(head *ListNode1) *ListNode1 {
	next := head.Next
	newOne := head
	newOne.Next = nil

	for next != nil{
		temp := newOne
		newOne = next
		next = next.Next
		newOne.Next = temp
	}
	return newOne
}


// 迭代2（正统）
func reverseList1(head *ListNode1) *ListNode1{
	if head == nil || head.Next == nil {     //链表为空或者仅1个数直接返回
		return head
	}
	curHead := head
	var newHead *ListNode1

	for curHead != nil{
		nextTemp := curHead.Next     //暂存h下一个地址，防止变化指针指向后找不到后续的数
		curHead.Next = newHead       //h->next指向前一个空间
		newHead = curHead            //新链表的头移动到h，扩长一步链表
		curHead = nextTemp           //h指向原始链表h指向的下一个空间
	}
	return newHead
}

// 递归
func reverseList2(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}

	newList := reverseList1(head.Next)  //一直循环到链尾, 在倒数第二个节点就能返回倒数第一个节点
	head.Next.Next = head               //使当前节点的后一个节点指向当前节点, 翻转链表的指向
	head.Next = nil                     //其实当前节点已经是处于新链表的末尾，置空当前节点的指向
	return newList                      //新链表头永远指向的是原链表的链尾
}