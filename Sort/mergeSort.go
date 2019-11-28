/*
归并排序：

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	l := 10
	arr := make([]int,l)
	for i := 0;i<l;i++{
		arr[i] = i
	}
	rand.Seed(time.Now().Unix())
	for i := len(arr) - 1;i>=0;i--{
		j := rand.Int() % (i + 1)
		arr [i],arr[j] = arr[j],arr [i]
	}
	fmt.Println(arr)
	mergeSort(arr,0,len(arr) - 1)
	fmt.Println(arr)
}

func mergeSort(arr []int, l int, r int){
	if l == r {
		return
	}
	m := l + int((r - l) / 2)
	mergeSort(arr, l, m)
	mergeSort(arr, m + 1 , r)
	mergeFunc(arr, l, m, r)
}

func mergeFunc(arr []int, l int, m int, r int){
	i := 0
	p1 := l
	p2 := m + 1
    newArr := make([]int,r - l + 1)
    for p1 <= m && p2 <= r{
		if arr[p1] < arr[p2]{
			newArr[i]=arr[p2]
			p2 ++
		}else{
			newArr[i]= arr[p1]
			p1 ++
		}
		i++
	}

	for p1 <= m {
		newArr[i]= arr[p1]
		i ++
		p1 ++
	}

	for p2 <= r {
		newArr[i]= arr[p2]
		i ++
		p2 ++
	}

	for j := 0;j< r - l + 1;j++{
		arr[l + j] = newArr[j]
	}

}