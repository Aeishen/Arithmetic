/*
   @File: CanPartitionKSubsets
   @Author: Aeishen
   @Date: 2019/12/30 9:58
   @Description:698. 划分为k个相等的子集
*/

/*
题目：
	给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。


示例 ：
	输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
	输出： True
	说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。

注意：
	1 <= k <= len(nums) <= 16
	0 < nums[i] < 10000

思路：
	递归：
		通过构造"子集和"进行搜索: 模拟构造 k 个子集（nums 的不相交子集）。对于 nums 中的每个数字，我们将其放入第 i 子集检查是否解决了问题。
        我们可以通过递归搜索来检查是否存在可能性。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{4, 3, 2, 3, 5, 2, 1}
	//k := 4
	//fmt.Println(canPartitionKSubsets(nums, k))
	//nums := []int{10,10,10,7,7,7,7,7,7,6,6,6}
	//k := 3
	//fmt.Println(canPartitionKSubsets(nums, k))
	nums := []int{960,3787,1951,5450,4813,752,1397,801,1990,1095,3643,8133,893,5306,8341,5246}
	k := 6
	fmt.Println(canPartitionKSubsets(nums, k))
}

var target int
func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)

	sum := 0
	for _,v := range nums{
		sum = sum + v
	}
	target = sum / k
    if k * target < sum || nums[len(nums) - 1] > target{
    	return false
	}

	l := len(nums)
	for nums[l - 1] == target {
		k --
		l --
	}

	return search(make([]int,k), nums, l)
}

func search(newGroups []int, nums []int, l int)bool{
	if l == 0 {
		return true
	}
    limitIndex := l - 1
	for i := 0; i < len(newGroups); i ++{
		temp := newGroups[i] + nums[limitIndex]
		if temp <= target{
			if temp < target && temp + nums[0] > target{
				continue
			}
			newGroups[i] = temp
			if search(newGroups, nums[:limitIndex], l - 1) {
				return true
			}
			newGroups[i] = temp - nums[limitIndex]
		}
	}
	return false
}