/*
   @File : stringInReverseOrder
   @Author: Aeishen
   @Date: 2019/12/16 21:26
   @Description: 以相反的顺序打印字符串
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