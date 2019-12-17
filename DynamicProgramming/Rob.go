/*
@author : Aeishen
@data :  19/10/05, 17:35
@description :
*/

/*
题目：
    你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
    如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情
    况下，能够偷窃到的最高金额。
示例 1:
	输入: [1,2,3,1]
	输出: 4
	解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
		 偷窃到的最高金额 = 1 + 3 = 4 。输入: [1,2,3,1]

示例 2:
	输入: [2,7,9,3,1,7]
	输出: 12
	解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 6 号房屋 (金额 = 7)。
		 偷窃到的最高金额 = 2 + 9 + 7 = 18 。
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rob_dp([]int{1,2,3,1}))
	fmt.Println(rob_dp([]int{2,7,9,3,1,7}))
}

func rob_dp(nums []int) int {
	if len(nums) <= 0{
		return 0
	}
	a,b := 0,0
	for i := 0; i < len(nums); i++{
		if i % 2 != 0{
			a = int(math.Max(float64(a + nums[i]),float64(b)))
		}else{
			b = int(math.Max(float64(b + nums[i]),float64(a)))
		}
	}
    fmt.Println(a,b)
	return int(math.Max(float64(a),float64(b)))
}