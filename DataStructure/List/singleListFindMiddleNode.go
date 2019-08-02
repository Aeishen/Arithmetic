/*
@author : Aeishen
@data :  19/07/25, 10:19

@description :
*/

/*
题目：找出单链表的中间结点（在这里不使用链表长度），给定链表头结点，在最小复杂度下输出该链表的中间结点

思路：如果只知链表的头结点，我们一般的思路就是先遍历链表得到链表长度，然后再遍历一遍得到中间结点，如此
	 时间复杂度为$O(n)+O(\frac n2)$。
     上面的思路似乎不太令人满意。我们又想到快慢指针，它有一个很重要的性质：慢指针走的长度等于快慢指针
     相距的程度。所以利用这个性质，当快指针走到链表尾时，慢指针正好在中间结点。
*/

package main

import (
	"Github/Arithmetic/DataStructure/List/list"
	"fmt"
)

func createList(arr []int) *List.MyList{
	l := List.InitList()
	for _,v:= range arr{
		l.Add(v)
	}
	return l
}

func findMiddle(l *List.MyList) *List.Node{
	_,slowPtr := l.GetHead()
	fastPtr := slowPtr

	for fastPtr.Next != nil && fastPtr.Next.Next != nil{
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	return slowPtr
}

func main() {
	arr := []int{1,4,9,16,25,36,49}
	l := createList(arr)
	fmt.Println(findMiddle(l).Data)
}