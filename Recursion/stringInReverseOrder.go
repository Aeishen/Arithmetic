/*
   @File : stringInReverseOrder
   @Author: Aeishen
   @Date: 2019/12/16 21:26
   @Description: 以相反的顺序打印字符串
*/

/*

思路：
    递归：
		首先，我们可以将所需的函数定义为 printReverse(str[0...n-1])，其中 str[n-1] 表示字符串中的最后一个字符。
        然后我们可以分两步完成给定的任务：
		printReverse(str[:len(str) -1])：以相反的顺序打印子字符串 str[0...n-2] 。
		print(str[len(str) -1])：打印字符串中的最后一个字符str[n-1]。
*/

package main

import "fmt"

func main() {
	printReverse("asdfg")
}

func printReverse(s string){
	if len(s) <= 0{
		return
	}
	fmt.Println(string(s[len(s) -1]))
	printReverse(s[:len(s) -1])
}