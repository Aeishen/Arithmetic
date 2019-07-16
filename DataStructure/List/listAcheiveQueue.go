/*
@author : Aeishen
@data :  19/07/15, 14:19

@description : 单链表实现队列
*/

package main

import (
	"./list"
	"fmt"
)

type myQueue struct {
	Items *List.MyList
}

//入队
func (q *myQueue)Enqueue(data List.Object)  {
	if q.Items.GetSize() == 0 {
		q.Items.Add(data)
	}else{
		q.Items.Insert(1,data)
	}
}

//出队
func (q *myQueue)Dequeue()  {
	if q.Items.GetSize() == 0 {
		return
	}else{
		q.Items.Delete(q.Items.GetSize())
	}
}

//获取队列首元素（最先进入队列的）
func (q *myQueue)GetFront() List.Object {
	if q.Items.GetSize() == 0 {
		return nil
	}else{
		_,tailNode := q.Items.GetTail()
		return tailNode.Data
	}
}

//获取队列
func NewQueue() *myQueue{
	s := &myQueue{List.InitList()}
	return s
}

func main() {
	//测试链表实现的队列
	s := NewQueue()
	for i:= 10;i>6;i--{
		s.Enqueue(i)
	}
	fmt.Printf("%v\n",s.Items.Traverse())
	fmt.Printf("%v\n",s.GetFront())
	s.Dequeue()
	fmt.Printf("%v\n",s.GetFront())
	s.Dequeue()
	fmt.Printf("%v\n",s.GetFront())
	s.Dequeue()
	fmt.Printf("%v\n",s.GetFront())
	s.Dequeue()
	fmt.Printf("%v\n",s.GetFront())
}