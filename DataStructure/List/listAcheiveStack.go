/*
@author : Aeishen
@data :  19/07/15, 14:19

@description : 单链表实现栈
*/

package main

import (
	"./list"
	"fmt"
)

type myStack struct {
	Items *List.MyList
}

//入栈
func (s *myStack)push(data List.Object)  {
	if s.Items.GetSize() == 0 {
		s.Items.Add(data)
	}else{
		s.Items.Insert(1,data)
	}
}

//出栈
func (s *myStack)pop()  {
	if s.Items.GetSize() == 0 {
		return
	}else{
		s.Items.Delete(1)
	}
}

//获取栈顶元素
func (s *myStack)peep() List.Object {
	if s.Items.GetSize() == 0 {
		return nil
	}else{
		_,headNode := s.Items.GetHead()
		return headNode.Data
	}
}

//获取栈
func NewStack() *myStack{
	s := &myStack{List.InitList()}
	return s
}

func main() {
	//测试链表实现的栈
	s := NewStack()
	for i:= 1;i<5;i++{
		s.push(i)
	}
	fmt.Printf("%v\n",s.Items.Traverse())
	fmt.Printf("%v\n",s.peep())
    s.pop()
	fmt.Printf("%v\n",s.peep())
	s.pop()
	fmt.Printf("%v\n",s.peep())
	s.pop()
	fmt.Printf("%v\n",s.peep())
}