/*
题目：
	定义栈的数据结构，请在该类型中实现一个能够得到栈最小元素的 min 函数。
*/

/*
解题思路：
	使用两个 stack，一个作为数据栈，另一个作为辅助栈。其中 数据栈 用于
    存储所有数据，而 辅助栈 用于存储最小值。
*/

package main

import (
	"./stack"
)

type newStack struct {
	dataStack *Stack.ItemStack
	minStack *Stack.ItemStack
}

func initNewStack() *newStack {
	ns := newStack{
		Stack.NewStack(),
		Stack.NewStack(),
	}
	return &ns
}


func (this *newStack)push(i Stack.Item){
	this.dataStack.Push(i)
	if this.minStack.IsEmpty(){
		this.minStack.Push(i)
	}else{
		if this.minStack.Peek().(int) >= i.(int){
			this.minStack.Push(i)
		}
	}
}

func (this *newStack)pop(){
	if this.dataStack.IsEmpty() || this.minStack.IsEmpty(){
		return
	}
	if this.dataStack.Peek() == this.minStack.Peek(){
		this.minStack.Pop()
	}
	this.dataStack.Pop()
}

func (this *newStack)peek() Stack.Item{
	if this.dataStack.IsEmpty() || this.minStack.IsEmpty(){
		return 0
	}
	return this.dataStack.Peek()
}

func (this *newStack)min() Stack.Item{
	return this.minStack.Peek()
}

func main() {
	//n := initNewStack()
	//arr := []int{3,2,4,1,5}
	//for _,v := range arr{
	//	n.push(v)
	//}
	//fmt.Println("所有元素入栈后：-----------")
	//fmt.Println("dataStack:",n.dataStack.Items)
	//fmt.Println("minStack:",n.minStack.Items)
	//fmt.Println("最小元素：",n.min())
	//fmt.Println("栈顶元素：",n.peek())
	//
	//fmt.Println("\n第一次出栈后：-----------")
	//n.pop()
	//fmt.Println("dataStack:",n.dataStack.Items)
	//fmt.Println("minStack:",n.minStack.Items)
	//fmt.Println("最小元素：",n.min())
	//fmt.Println("栈顶元素：",n.peek())
	//
	//fmt.Println("\n第二次出栈后：-----------")
	//n.pop()
	//fmt.Println("dataStack:",n.dataStack.Items)
	//fmt.Println("minStack:",n.minStack.Items)
	//fmt.Println("最小元素：",n.min())
	//fmt.Println("栈顶元素：",n.peek())
}

