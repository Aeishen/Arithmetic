//快速排序
/*
快速排序和递归排序都是分治法，不同点在于递归排序是直接将数组平均分组，而快速排序是以某个值v为基准，
分组后，左边的值小于v，右边的值大于v，等于的情况放左边右边都可以
*/


package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 选第一个元素为基数的单路快排1
func partition_1(arr []int, startPos int, endPos int) int {
	i,j, base:= startPos,endPos,arr[startPos]
	for i < j{
		for i < j{
			if arr[j] > base{
				arr[i] = arr[j]
				i ++
				break
			}
			j --
		}
		for i < j{
			if arr[i] < base{
				arr[j] = arr[i]
				j --
				break
			}
			i ++
		}
	}
	arr[i]  = base
	return i
}

// 选第一个元素为基数的单路快排2
func partition_2(arr []int,start ,end int)(p int){
	i,j,base := start,end ,arr[start]

	for i<j  {
		// 获取比base小的元素的索引
		for i < j && arr[j] > base  {
			j--
		}
		// 获取比base大的元素的索引
		for i < j && arr[i] <= base  {
			i++
		}
		if i < j {
			arr[i] ,arr[j] = arr[j],arr[i]
		}
	}
	arr[i],arr[start] = arr[start],arr[i]
	return i
}

// 选第一个元素为基数的单路快排3
func partition_3(arr []int,start ,end int)(p int){
	i := start
    base := arr[start]
	for j := i + 1;j <= end;j++{
		if arr[j] < base{
			i ++
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	arr[i],arr[start] = arr[start],arr[i]
	return i
}


// 选随机为基数的单路快排1
func partition_4(arr []int,start,end int)(p int){
	i := start
	randIndex := start + rand.Intn(end)
	base := arr[randIndex]
	for j := i + 1;j <= end;j++{
		if arr[j] < base{
			i ++
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	arr[i],arr[start] = arr[start],arr[i]
	return i
}

// 选第一个元素为基数的单路快排
func quickSort1(arr []int, startPos int, endPos int)  {
	if startPos >= endPos{
		return
	}
	//pos := partition_1(arr, startPos, endPos)
	//pos := partition_2(arr, startPos, endPos)
	pos := partition_3(arr, startPos, endPos)
	quickSort1(arr, startPos, pos - 1)
	quickSort1(arr, pos + 1, endPos)
}

// 选随机元素为基数的单路快排
func quickSort2(arr []int, startPos int, endPos int)  {
	if startPos >= endPos{
		return
	}
	pos := partition_4(arr, startPos, endPos)
	quickSort1(arr, startPos, pos - 1)
	quickSort1(arr, pos + 1, endPos)
}



func main() {
	arr1 := getRandomArray(5000,1000)  //无序数组
	arr2 := getSortArray(5000)                    //有序数组
	arr3 := getSortArray(5000)                    //有序数组

	t1 := time.Now()
	fmt.Println("The randomArr before sort is:", arr1)
	quickSort1(arr1,0,len(arr1)-1)
	fmt.Println("The randomArr after sort is:", arr1,t1.Sub(time.Now()))

	t2 := time.Now()
	fmt.Println("The sortArr before sort is:", arr2)
	quickSort1(arr2,0,len(arr2)-1)
	fmt.Println("The sortArr after sort is:", arr2,t2.Sub(time.Now()))

	t3 := time.Now()
	fmt.Println("The sortArr before sort is:", arr3)
	quickSort2(arr3,0,len(arr3)-1)
	fmt.Println("The sortArr after sort is:", arr3,t3.Sub(time.Now()))

	/*
		当对有序数组使用快速排序时，每次都取数组首个值为基准，这样分区后左右数组始终不平衡，
	    在最坏的情况下，快排的时间复杂度会退化成O(n^2)，此时每次随机取值为基准，能避免快
	    速排序往最坏的情况发展，加快排序时间

	    但是当数组中存在大量重复元素的时候，随机快排的速度反而不如基本快排，这是因为重复元素
		会偏向左边或右边，会导致分区不平均，拖慢排序时间，此时要使用双路快排或者三路快排
	*/
}


// 获取随机数组
func getRandomArray(numCount int,rangeCount int)(arr []int){
	for i := 0;i < numCount; i++{
		arr = append(arr,rand.Intn(rangeCount))
	}
	return
}

// 获取有序数组
func getSortArray(numCount int)(arr []int){
	for i := 0;i < numCount; i++{
		arr = append(arr,i)
	}
	return
}

