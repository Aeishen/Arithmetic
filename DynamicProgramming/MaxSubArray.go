/*
@author : Aeishen
@data :  19/10/09, 22:13
@description :
*/

/*
题目：
    给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
	输入: [-2,1,-3,4,-1,2,1,-5,4],
	输出: 6
	解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

进阶:
	如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

思路:
*/
package main

import "fmt"

func main() {
	fmt.Println(maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray1([]int{-2}))
}

func maxSubArray(nums []int) int {
	maxValue := nums[0]  // 没加上当前元素前的最大和
	curValue := 0        // 加上当前元素后的和
	for i := 0;i<len(nums);i++{
		if curValue > 0{
			curValue = curValue + nums[i]
		}

		if curValue <= 0{
			curValue = nums[i]
		}

		if curValue > maxValue{
			maxValue = curValue
		}
	}
	return maxValue
}

// 如果规定和为负数时返回0，可以这样写
func maxSubArray1(nums []int) int {
	maxValue := nums[0]  // 没加上当前元素前的最大和
	curValue := 0        // 加上当前元素后的和
	for i := 0;i<len(nums);i++{
		if curValue + nums[i] > 0{
			curValue = curValue + nums[i]
		}else{
			curValue = 0
		}
		if curValue > maxValue{
			maxValue = curValue
		}
	}
	return maxValue
}