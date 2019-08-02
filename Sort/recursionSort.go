/*
@author : Aeishen
@data :  19/07/25, 11:58

@description :递归排序
*/
package main

import "fmt"

func merger(arr1 []int,arr2 []int)(mergerArr []int){
	len_arr1 := len(arr1)
	len_arr2 := len(arr2)
	i,j := 0,0
	for i < len_arr1 || j < len_arr2{
		if j == len_arr2{
			mergerArr = append(mergerArr,arr1[i])
			i++
		}else if i == len_arr1{
			mergerArr = append(mergerArr,arr2[j])
			j++
		}else if arr1[i] < arr2[j]{
			mergerArr = append(mergerArr,arr1[i])
			i++
		}else {
			mergerArr = append(mergerArr,arr2[j])
			j++
		}
	}
	return
}

// 自上而下
func recursionSort_1(arr []int) []int{
	len_arr := len(arr)
	if len_arr == 1{
		return arr
	}
	mid := len_arr / 2
	return merger(recursionSort_1(arr[:mid]),recursionSort_1(arr[mid:]))
}


func main() {
	arr := []int{12,3,4,57,1,5}
	fmt.Println(recursionSort_1(arr))
}