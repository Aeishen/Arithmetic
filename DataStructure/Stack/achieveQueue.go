/*
题目：
	用两个栈来实现一个队列，完成队列的 Push 和 Pop 操作。
*/

/*
解题思路：
	in 栈用来处理入栈（push）操作，out 栈用来处理出栈（pop）操作。一个元素进入 in 栈之后，出栈的顺序被反转。
    当元素要出栈时，需要先进入 out 栈，此时元素出栈顺序再一次被反转，因此出栈顺序就和最开始入栈顺序是相同的，
    先进入的元素先退出，这就是队列的顺序。
*/

package main

import (
	"./stack"
	"fmt"
)

//入队
func push(inStack *Stack.ItemStack,i Stack.Item){
	inStack.Items = append(inStack.Items,i)
}

//出队
func pop(outStack *Stack.ItemStack,inStack *Stack.ItemStack) (i Stack.Item){
	for !inStack.IsEmpty(){
		outStack.Push(inStack.Peek())
		inStack.Pop()
	}
	if outStack.IsEmpty(){
		return nil
	}
	i = outStack.Peek()
	outStack.Pop()
	return
}

func main() {
	inStack := Stack.NewStack()
	outStack := Stack.NewStack()

	for i := 1;i < 6; i++{
		push(inStack,i)
	}
	fmt.Println(inStack)

    for {
		i := pop(outStack,inStack)
		fmt.Println(outStack,i)
		if outStack.IsEmpty(){
			break
		}
	}
}