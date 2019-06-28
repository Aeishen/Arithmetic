// 冒泡排序
/*
从第一个元素开始，相邻的两个元素进行比较，小（大）的元素往后移，第一趟下来整个数组最小（最大）的元素一定是位于数组末位
往后类推
*/

package main

import "fmt"

func bubbleSort(arr []int, )  {
	len_arr := len(arr)

	// i 在这里可以表示为已经排好序的元素个数，也可表示为进行第几趟从第 j 个元素到最后一个未排序的元素之间比较
	for i := 0; i < len_arr; i++{

		// 从第 j 个元素到最后一个未排序的元素之间比较，使得 ”所有未排序的元素“ 中最小（最大）的元素移动到 ”所有未排序的元素“ 的末位，
		// 完成 i 趟从第 j 个元素到最后一个未排序的元素之间比较后，则未排序的元素个数就减少为（所有元素 - i）位
		// 继续排序剩余的未排序的元素

		for j := 0; j < (len_arr - 1 - i); j++{
			if arr[j] < arr[j + 1]{
				arr[j] , arr[j + 1] = arr[j + 1] , arr[j]
			}
		}
	}
}

func main() {
	arr := []int{12,3,4,57,1,5}

	fmt.Println("The arr before sort is:", arr)
	bubbleSort(arr)
	fmt.Println("The arr after sort is:", arr)
}
