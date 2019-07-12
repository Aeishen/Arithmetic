/*
题目：
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效，返回true/false。
	有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。
*/

/*
解题思路：
	遍历输入字符串
	如果当前字符为左半边括号时，则将其压入栈中
	如果遇到右半边括号时，分类讨论：
	1）如栈不为空且为对应的左半边括号，则取出栈顶元素，继续循环
	2）若此时栈为空，则直接返回false
	3）若不为对应的左半边括号，反之返回false
*/
package main

import (
	"./stack"
	"fmt"
)

func check(chars []byte) bool{
	s := Stack.InitStack()
	len_c := len(chars)
	for i,v := range chars{
		if v == '(' || v == '{' || v == '['{
			fmt.Printf("[]byte入栈元素的动态类型：%T,动态值：%#v\n",v,v)
			s.Push(v)
		}else{
			if s.IsEmpty() {
				return false
			}else{
				fmt.Printf("[]byte出栈元素的动态类型：%T,动态值：%#v\n",s.Peek(),s.Peek())
				if (v == ')' && s.Peek().(byte) == '(')||(v == '}' && s.Peek().(byte) == '{')||(v == ']' && s.Peek().(byte) == '['){
					s.Pop()
				}
				if i ==  len_c - 1{
					return s.IsEmpty()
				}
			}
		}
	}
	return false
}

func checks(chars string) bool{
	s := Stack.InitStack()
	len_c := len(chars)
	for i,v := range chars{
		if v == '(' || v == '{' || v == '['{
			fmt.Printf("string入栈元素的动态类型：%T,动态值：%#v\n",v,v)
			s.Push(v)
		}else{
			if s.IsEmpty() {
				return false
			}else{
				fmt.Printf("string出栈元素的动态类型：%T,动态值：%#v\n",s.Peek(),s.Peek())
				if (v == ')' && s.Peek() == '(')||(v == '}' && s.Peek() == '{')||(v == ']' && s.Peek() == '['){
					s.Pop()
				}
				if i ==  len_c - 1{
					return s.IsEmpty()
				}
			}
		}
	}
	return false
}

func main() {

	// 使用字符切片
	var chars []byte
	//chars = []byte{'(',')','{','}','[',']'}
	//chars = []byte{'(',')','{','}','[','}'}
	//chars = []byte{'(','[',')',']'}
	//chars = []byte{'(','{','}',')'}
	chars = []byte{'(',')'}
	//chars = []byte{'(',']'}
	//chars = []byte{')','(','[',']'}
	isValid := check(chars)
	fmt.Println(isValid)

	//使用字符串
	var str string
	//str = "()[]{}"
	//str = "()"
	str = "([)]"
	isValids := checks(str)
	fmt.Println(isValids)

}