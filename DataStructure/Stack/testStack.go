package main

import (
	"./stack"
	"fmt"
)


func main() {
	test_stack := Stack.NewStack() //获取栈

	fmt.Println("入栈前栈是否为空：",test_stack.IsEmpty()) //栈判空

	for i:= 1;i<11;i++{
		test_stack.Push(i*i)          //入栈
	}

	fmt.Println("入栈后栈是否为空：",test_stack.IsEmpty()) //栈判空

	fmt.Println("入栈后栈元素个数：",test_stack.GetSize())//获取栈元素个数

	fmt.Println("入栈后栈顶元素是：",test_stack.Peek())    // 获取栈顶元素

	test_stack.Pop()                // 出栈

	fmt.Println("出栈后栈元素个数：",test_stack.GetSize())//获取栈元素个数

	fmt.Println("出栈后栈顶元素是：",test_stack.Peek())    // 获取栈顶元素


}