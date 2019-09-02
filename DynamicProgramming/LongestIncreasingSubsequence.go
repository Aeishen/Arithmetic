/*
@author : Aeishen
@data :  19/08/09, 15:42

@description : LeetCode 上第 300 号问题:最长上升子序列
*/

/*
题目：
	给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:
	输入: [10,9,2,5,3,7,101,18]
	输出: 4
	解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:
	可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
	你算法的时间复杂度应该为 O(n2) 。

进阶:
	你能将算法的时间复杂度降低到 O(n log n) 吗?
*/


/*
思路：
	动态规划：时间复杂度 O(N^2)
            1.定义一个dp切片，dp[i]用来表达nums[i]这个数结尾的最长递增子序列的长度.
            2.那么最终结果就是取dp中的最大值
			3.关键是获取dp[i]的值，dp[i]的值就是前面那些结尾比nums[i]小的子序列中最长的那个+1
    二分查找法：时间复杂度  O(NlogN)

*/

package main

import "fmt"

// 动态规划
func lengthOfLIS_dp(nums []int) int{
	l := len(nums)
	longest := 0
	if l <= 0{
		return longest
	}

    // 初始化dp，因为子序列最少也要包含自己，所以长度最小为 1
	dp := make([]int,l)
    dp[0] = 1

    // 设置dp
	for i := 0; i< l; i++{
		dp[i] = getDpValue(nums , dp, i)
		if longest < dp[i]{
			longest =  dp[i]
		}
	}
	return longest
}

//获取dp[i]的值
func getDpValue(nums []int, dp []int, i int) int{
	curLongest := 0
	for j := 0; j<i ; j++{
		if nums[j] < nums[i] && dp[j] > curLongest {
			curLongest = dp[j]
		}
	}
	return curLongest + 1
}

// 二分查找法
func lengthOfLIS_bs(nums []int) int{
	l := len(nums)
	piles := 0                 //堆数
    top := make(map[int]int)   //每一堆牌顶部的那张牌
	for i:= 0;i<l;i++ {
		curNum := nums[i]

		//搜索左侧边界的二分查找
		left,right := 0, piles
		for left < right{
			mid := left + (right - left)/2
			if curNum > top[mid]{
				left = mid + 1
			}else if curNum < top[mid] {
				right = mid - 1
			}else{
				right = mid
			}

		}
		if left == piles{
			piles ++
		}
		top[left] = curNum
	}

	return piles
}


func main() {
	fmt.Println(lengthOfLIS_dp([]int{10,9,2,5,3,7,101,18}) )
	fmt.Println(lengthOfLIS_dp([]int{}) )
	fmt.Println(lengthOfLIS_bs([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS_bs([]int{}) )
}
