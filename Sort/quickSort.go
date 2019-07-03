//快速排序
/*
选择一个基数，比它小的分为一部分，大于等于它的分为另外一部分
*/


package main

import "fmt"

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


func partition_2(arr []int,start ,end int)(p int)  {
	i,j,base := start,end ,arr[start]

	for i<j  {
		for i < j && arr[j] >base  {
			j--
		}
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

func quickSort(arr []int, startPos int, endPos int)  {
	if startPos >= endPos{
		return
	}
	i := partition_2(arr, startPos, endPos)
	quickSort(arr[:i], startPos, i - 1)
	quickSort(arr[i + 1:], i + 1, len(arr[i + 1:]) - 1)
}

func quick1(arr []int, s int,e int){
	if s > e {
		return
	}
	i := s
	j := e
	base := arr[s]

	for i<j{
		for i<j{
			if arr[j] < base{
				arr[i] = arr[j]
				i ++
				break
			}
			j --
		}
		for i<j{
			if arr[i] > base{
				arr[j] = arr[i]
				j --
				break
			}
			i ++
		}
	}
	arr[i] = base
	quick1(arr[:i], 0,len(arr[:i]) - 1)
	quick1(arr[i + 1:], i + 1,len(arr[i + 1:]) - 1)
}

func quick2(arr []int, s int,e int){
	if s > e {
		return
	}
	i := s
	j := e
	base := arr[s]
	for i<j{
		for i<j && arr[j] < base{
			j --
		}
		for i<j && arr[i] >= base{
			i ++
		}
		if i<j{
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	arr[i],arr[s] = arr[s],arr[i]
	quick2(arr[:i], 0,len(arr[:i]) - 1)
	quick2(arr[i + 1:], i + 1,len(arr[i + 1:]) - 1)
}

func main() {
	arr := []int{12,43,30,4,57,1,5}
	fmt.Println("The arr before sort is:", arr)
	quick2(arr,0,len(arr)-1)
	fmt.Println("The arr after sort is:", arr)
}
