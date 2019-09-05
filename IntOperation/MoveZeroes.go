/*
@author : Aeishen
@data :  19/09/02, 14:29
@description :
*/

/*
题目：
    给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例 1:
	输入: [0,1,0,3,12]
	输出: [1,3,12,0,0]

说明:
	必须在原数组上操作，不能拷贝额外的数组。
	尽量减少操作次数。

思路1：使用两个索引endIndex和curIndex，curIndex从后向前遍历直到其元素为0，然后根据lastIndex与curIndex之间的差值，将元素前移。
思路2：非0数前面有几个0就往前面移动几个位置
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{2,0,1,0,3,12}
	moveZeroes1(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int)  {
	curIndex := len(nums) - 1
	endIndex := curIndex

	for curIndex >= 0{
		if (nums[curIndex] == 0){
			count := endIndex - curIndex
			for i := 0;i < count; i ++{
				nums[curIndex + i]  = nums[curIndex + i + 1]
			}
			nums[endIndex] = 0
			endIndex --
		}
		curIndex --
	}
}

func moveZeroes1(nums []int)  {
	count := 0
	for i := 0;i<len(nums);i++{
		//if nums[i] == 0{
		//	count ++
		//}else if count > 0{
		//	nums[i - count] = nums[i]
		//	nums[i] = 0
		//}
		if nums[i] != 0 {
			nums[count],nums[i] = nums[i],nums[count]
			count ++
		}
	}
}
