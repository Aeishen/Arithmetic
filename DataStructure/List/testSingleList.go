/*
@author : Aeishen
@data :  19/07/15, 16:11

@description : 测试单链表
*/

package main

import (
	"./list"
	"fmt"
)

func main() {
	list := List.InitList()

	// 测试list.Add()
	for i:= 1;i<=10;i++{
		list.Add(i*i)
	}

	// 测试list.Get()
	if ok,data  := list.Get(11); ok{
		fmt.Println("获取链表第11个元素：",data)
	}else{
		fmt.Println("链表不存在第11个元素")
	}

	if ok,data  := list.Get(10); ok{
		fmt.Println("获取链表第10个元素：",data)
	}else{
		fmt.Println("链表不存在第10个元素")
	}

	if ok,data  := list.Get(5); ok{
		fmt.Println("获取链表第5个元素：",data)
	}else{
		fmt.Println("链表不存在第5个元素")
	}

	if ok,data  := list.Get(1); ok{
		fmt.Println("获取链表第1个元素：",data)
	}else{
		fmt.Println("链表不存在第1个元素")
	}

	// 测试list.Insert(),list.Traverse(),list.GetSize()
	fmt.Printf("在第9个元素后插入新元素99前遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())
    list.Insert(9,99)
	fmt.Printf("在第9个元素后插入新元素99后遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())
	list.Insert(list.GetSize(),99)
	fmt.Printf("在最后1个元素后插入新元素99后遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())


	// 测试list.Delete
	list.Delete(5)
	fmt.Printf("删除第5个元素遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())


	// 测试list.Search
	if ok,Indexs := list.Search(99);ok{
		fmt.Printf("查询链表中99元素的位置：%v,\n",Indexs)
	}
	if ok,Indexs := list.Search(49);ok{
		fmt.Printf("查询链表中49元素的位置：%v,\n",Indexs)
	}

	// 测试list.GetHead(),list.GetTail()
	if ok1,headNode := list.GetHead();ok1{
		fmt.Printf("获取头指针元素：%v\n",headNode.Data)
	}
	if ok2,tailNode := list.GetTail();ok2{
		fmt.Printf("获取尾指针元素：%v\n",tailNode.Data)
	}


	//测试list.RemoveAll()
	fmt.Printf("清空链表前遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())
	list.RemoveAll()
	fmt.Printf("清空链表后遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())


	// 测试list.Reversal()
	for i:= 1;i<=10;i++{
		list.Add(i*i)
	}
	fmt.Printf("反转前遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())
    list.Reversal()
	fmt.Printf("反转后遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())

	// 测试list.GetHead(),list.GetTail()
	if ok1,headNode := list.GetHead();ok1{
		fmt.Printf("反转后获取头指针元素：%v\n",headNode.Data)
	}
	if ok2,tailNode := list.GetTail();ok2{
		fmt.Printf("反转后获取尾指针元素：%v\n",tailNode.Data)
	}

	// 测试list.Insert(),list.Traverse(),list.GetSize()
	list.Insert(9,99)
	fmt.Printf("反转后在第9个元素后插入新元素99后遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())
	list.Insert(1,99)
	list.Set(5,1000)
	fmt.Printf("反转后在最后1个元素后插入新元素99后遍历链表：%v,链表长度：%v\n",list.Traverse(),list.GetSize())




}