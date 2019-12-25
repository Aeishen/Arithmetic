/*
   @File: MergeTwoLists
   @Author: Aeishen
   @Date: 2019/12/25 15:58
   @Description:
*/

/*
题目：
	将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例:
	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4

思路：
	递归：我们可以递归地定义在两个链表上进行合并（merge）操作的结果，如下所示（在不考虑空列表的情况下）：
			当list1[0] < list2[0]：  list1[0] + merge(list1[1:], list2)
			当list1[0] >= list2[0]： list2[0] + merge(list1, list2[1:])
	​     也就是说，我们取两个列表头部中较小的那个，然后再加上合并其余元素所得到的结果。我们直接对上述递归建模，首先考虑边界情况。
          具体来说，如果 l1 或 l2 最初为 nil，则不需要执行合并，我们只需返回非空列表。否则，我们需要确定 l1 和 l2 中哪个的头节点更小，
          并递归地处理该头节点的 next 值以得到下一次合并结果。 如果两个列表都以空结束，那么递归就会停止。
			时间复杂度：O(n+m)。因为每次递归调用都会将指向 l1 或 l2 的指针递增一次（逐渐接近每个列表末尾的 nil），所以每个列表中的
					每个元素都会对 mergeTwoLists 进行一次调用。因此，时间复杂度与两个列表的大小之和是线性相关的。
			空间复杂度：O(n+m)。一旦调用 mergetwolist，直到到达 l1 或 l2 的末尾时才会返回，因此 n + m 的栈将会消耗 O(n+m) 的空间。

    迭代：设置一个 prehead 节点（虚节点），这会帮助我们轻松地返回合并之后的列表的头节点。 我们还需要维护一个 prev 指针，它指向可能需
          要调整 next 指针的当前节点。 然后执行以下操作，直到 l1 和 l2 中至少有一个指向 nil：如果 l1 处的值小于或等于 l2 处的值，
          那么我们将 l1 连接到前一个节点，并递增 l1。 否则，我们对 l2 做同样的事情。 不管我们连接的是哪个列表，我们都会增加 prev，
          。循环终止后，l1 和 l2 中最多有一个是非空的。 因此（因为输入列表是按有序的），如果其中一个列表是非空的，那么它包含的元素
          一定大于所有先前合并的元素。 这意味着我们可以直接将非空列表连接到已合并列表并返回它。
			时间复杂度：O(n+m)。在每次循环迭代中， l1 和 l2 其中之一会递增，所以 for 循环运行的迭代次数等于这两个列表的长度之和。
                    所有其他的工作都不变，所以整体的复杂性是线性的。
			空间复杂度：O(1)。迭代方法只会分配几个指针，因此它的总体内存占用是恒定的。
*/

package main

import (
	"fmt"
)

type ListNode2 struct {
     Val int
     Next *ListNode2
}

func main() {
	l1 := &ListNode2{1,&ListNode2{2,&ListNode2{4,nil}}}
	l2 := &ListNode2{1,&ListNode2{3,&ListNode2{4,nil}}}

	fmt.Println(mergeTwoLists1(l1, l2))
}

// 递归
func mergeTwoLists1(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	if l1.Val < l2.Val{
		l1.Next = mergeTwoLists1(l1.Next,l2)
		return l1
	}else{
		l2.Next = mergeTwoLists1(l2.Next,l1)
		return l2
	}
}


// 递归
func mergeTwoLists2(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	prehead  := &ListNode2{-1,nil}
	prev := prehead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			prev.Next = l1
			l1 = l1.Next
		}else{
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil{
		prev.Next = l2
	}else{
		prev.Next = l1
	}

	return prehead.Next
}

