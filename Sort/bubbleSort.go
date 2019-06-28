// 冒泡排序
/*
从第一个元素开始，相邻的两个元素进行比较，小（大）的元素往后移，第一趟下来整个数组最小（最大）的元素一定是位于数组末位
往后类推
*/

package main

import "fmt"


func bubbleSort_WithOutOptimize(arr []int)  {
	len_arr := len(arr)

	sortCount := 0

	//i 在这里可以表示为已经排好序的元素个数，也可表示为进行第几趟从第 j 个元素到最后一个未排序的元素之间比较
	for i := 0; i < len_arr; i++{

		/*
			从第 j 个元素到最后一个未排序的元素之间比较，使得 ”所有未排序的元素“ 中最小（最大）的元素移动到
			”所有未排序的元素“ 的末位，完成 i 趟从第 j 个元素到最后一个未排序的元素之间比较后，
			则未排序的元素个数就减少为（所有元素 - i）位继续排序剩余的未排序的元素
		*/

		for j := 0; j < (len_arr - 1 - i); j++{
			sortCount ++
			if arr[j] < arr[j + 1]{
				arr[j] , arr[j + 1] = arr[j + 1] , arr[j]
			}
		}
	}
	fmt.Println("func bubbleSort_WithOutOptimize sortCount is :",sortCount)
}

/*
假如一个数组已经是有序的数组，则对其进行bubbleSort(arr)，那么bubbleSort内部同样会对该数组进行嵌套遍历，
浪费时间，下面为优化版本
*/

func bubbleSort_WithOptimize(arr []int)  {
	len_arr := len(arr)

	sortCount := 0

	for i := 0; i < len_arr; i++{

		// 定义一个标识,假设该数组已经有序了
		isSorted := true

		for j := 0; j < (len_arr - 1 - i); j++{
			sortCount ++
			if arr[j] < arr[j + 1]{
				arr[j] , arr[j + 1] = arr[j + 1] , arr[j]

				// 一旦发生交换说明前面假设不成立，该数组不是有序数组
				isSorted = false
			}
		}

		//一趟比较过后如果没有元素发生交换，前面假设成立，说明该数组已经是有序的，直接跳出循环
		if isSorted {
			break
		}
	}

	fmt.Println("func bubbleSort_WithOutOptimize sortCount is :",sortCount)
}


func main() {
	arr := []int{12,45,36,27,18,15}

	fmt.Println("The arr before bubbleSort_WithOutOptimize is:", arr)
	bubbleSort_WithOutOptimize(arr)
	fmt.Println("The arr after bubbleSort_WithOutOptimize is:", arr)


	arr = []int{12,45,36,27,18,15}

	fmt.Println("The arr before bubbleSort_WithOptimize is:", arr)
	bubbleSort_WithOptimize(arr)
	fmt.Println("The arr after bubbleSort_WithOptimize is:", arr)
}
