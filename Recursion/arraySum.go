/*
@author : Aeishen
@data :  19/07/19, 10:08

@description : 递归求数组之和
*/

package main

import "fmt"

func getSum(arr []int) int {
	l := len(arr)
	if l == 1 {
		return arr[0]
	}else{
		return getSum(arr[:l-1]) + arr[l - 1]
	}
}

func main() {
	arr := []int{8,2,6,4,8,2,6,4,5,9}
	fmt.Println(getSum(arr))
}