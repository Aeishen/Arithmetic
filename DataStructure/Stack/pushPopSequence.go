/*
题目：
	输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
    例如序列 1，2，3，4，5 是某栈的压入顺序，序列 4，5，3，2，1是该压栈序列对应的一个弹出序列，但4，3，5，1，2就
    不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）。
*/

/*
解题思路：
	借用一个辅助的栈，遍历压栈顺序，先讲第一个放入栈中，这里是 1，然后判断栈顶元素是不是出栈顺序的第一个元素，这里是
    4，很显然 1≠4 ，所以需要继续压栈，直到相等以后开始出栈。出栈一个元素，则将出栈顺序向后移动一位，直到不相等，这样
    循环等压栈顺序遍历完成，如果辅助栈还不为空，说明弹出序列不是该栈的弹出顺序。
*/


package main

import (
	"./stack"
	"fmt"
)

func isPopOrder(pushSequence []int,popSequence []int) bool{
	len_push := len(pushSequence)
	len_pop := len(popSequence)
	s := Stack.InitStack()

	for pushIndex,popIndex := 0,0; pushIndex < len_push; pushIndex++ {
		s.Push(pushSequence[pushIndex])
		for !s.IsEmpty() && popIndex < len_pop && s.Peek() == popSequence[popIndex] {
			s.Pop()
			popIndex ++
		}
	}
	return s.IsEmpty()

}
func main() {
	pushSequence := []int{1,2,3,4,5}
	//popSequence := []int{3,4,5,2,1}
	//popSequence := []int{2,3,4,5,1}
	//popSequence := []int{1,2,3,4,5}
	//popSequence := []int{4,5,3,2,1}
	//popSequence := []int{4,3,5,1,2}
	popSequence := []int{5,4,3,2,1}
	fmt.Println(isPopOrder(pushSequence,popSequence))
}