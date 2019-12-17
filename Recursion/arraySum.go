/*
@author : Aeishen
@data :  19/07/19, 10:08

@description : 递归求数组之和
*/

package main

import "fmt"


func getSum1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return getSum1(arr[1:]) + arr[0]  //不是尾部递归，因为它在递归调用返回后做一些计算。
}


func getSum2(arr []int) int {
	return helper_getSum2(arr, 0)
}

func helper_getSum2(a []int, s int) int {
	if len(a) == 0 {
		return s
	}
	return helper_getSum2(a[1:],a[0] + s) //这是尾部递归，因为最后一条指令是递归调用。
}


func main() {
	arr := []int{8,2,6,4,8,2,6,4,5,9}
	fmt.Println(getSum1(arr))
}