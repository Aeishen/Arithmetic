//插入排序
/*
1.所谓插入，就像打牌时候发完牌，将牌一张张按照一定的顺序插入到手中的牌。
从取到第一张牌开始，将第一张牌与剩余的牌假设性的为“已排好序的牌组”与“已排好序的牌组”

2.则从第2（n）张开始，它将插入到前面的有序牌组中，（假如排序顺序是从大到小）则它的位置应该是在所有比它小的牌的前面，
当它插入到该位置时，有序数组中比他小的牌的位置都会往后移动一个位置，此处的有序数组只是假设性的，

3.重复2，直到整个牌组有序
*/


package main

import "fmt"

func insertSort(arr []int){

	// i 是未排序元素的第一个元素的索引
	for i := 1; i < len(arr) ; i++{

		// 让未排序元素的第一个元素与前面已经排好序的元素（从最后一个元素开始）比较，找到比它小的元素则交换位置，直到没有元素比它小
		for j := i; j > 0; j--{
			if arr[j] > arr[j - 1]{
				arr[j] , arr[j - 1] = arr[j - 1] , arr[j]
			}
		}
	}
}

func main() {
	arr := []int{12,3,4,57,1,5}
	fmt.Println("The arr before sort is:", arr)
	insertSort(arr)
	fmt.Println("The arr after sort is:", arr)
}
