//选择排序
/*
1.假设将数组分为两个部分，前一个部分是选择出来已经排好序的，另一部分是还未被选择的元素集

2.每次从还未被选择的元素集选择最大（最小）的元素插入到已经排好序部分的末位

3.直到最后还未被选择的元素集只剩一个元素，则直接插入到已经排好序部分的末位
*/


package main

import "fmt"

func selectSort(arr []int){

	// i 可以理解为未被选择的元素集的第一个元素
	for i := 0; i < len(arr) - 1 ; i++{

		 //假设未被选择的元素集的第一个元素是未被选择的元素集中最大的元素
		 maxIndex := i

		// 未被选择的元素集相邻元素两两比较获取最大元素的索引
		for j := i + 1; j < len(arr); j++{
			if arr[j] > arr[maxIndex]{
				maxIndex = j
			}
		}

		// 如果未被选择的元素集中最大的元素不是元素集的第一个元素，则将则两个元素交换,
		// 这样就使得数组前面的i + 1（因为下标是从0开始的）个元素都是排好序的
		if maxIndex != i{
			arr[maxIndex],arr[i] = arr[i],arr[maxIndex]
		}

	}
}

func main() {
	arr := []int{12,3,4,57,1,5}
	fmt.Println("The arr before sort is:", arr)
	selectSort(arr)
	fmt.Println("The arr after sort is:", arr)
}
