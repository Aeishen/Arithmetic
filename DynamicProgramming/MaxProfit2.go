/*
@author : Aeishen
@data :  19/10/12, 23:42
@description :
*/

/*
题目：
    定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。如果你最多只允许完成两笔交易，规则是必须一笔成交后进行另一笔(即买-卖-买-卖的顺序进行)）
    设计一个算法来计算你所能获取的最大利润。

示例 1:
	输入: [20,10,22,15,75,37,80,50]
	输出: 87
	解释: 在第 2 天（股票价格 = 10）的时候买入，在第 5 天（股票价格 = 75）的时候卖出，最大利润 = 65，
         在第 6 天（股票价格 = 37）的时候买入，在第 7 天（股票价格 = 80）的时候卖出，最大利润 = 43
		 最大利润为 43 + 65 = 108

注意：对比GreedyAlgorithm/MaxProfit3.go、DynamicPrograming/MaxProfit1.go

思路: 1
     2：
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit2([]int{20,10,22,15,75,37,80,50}))
	fmt.Println(maxProfit2([]int{10,22,5,75,65,80}))
	fmt.Println(maxProfit2([]int{10,22,33,75,80}))
}

func maxProfit2(prices []int) int {
	l := len(prices)
	if l <= 1{
		return 0
	}
	endProfit := 0
	fromLeftProfits := make([]int,l)
	fromRightProfits := make([]int,l)

	minV := prices[0]
	maxV := prices[l-1]

	for i:= 1;i<l;i++{
		minV = int(math.Min(float64(minV), float64(prices[i])))
		fromLeftProfits[i] = int(math.Max(float64(fromLeftProfits[i]), float64(prices[i] - minV)))
	}

	for i:= l - 2;i >= 0;i--{
		maxV = int(math.Max(float64(maxV), float64(prices[i])))
		fromRightProfits[i] = int(math.Max(float64(fromRightProfits[i + 1]), float64(maxV - prices[i])))
	}
	fmt.Println(fromRightProfits,fromLeftProfits)
	for i := 0;i<l;i++{
		//endProfit = int(math.Max(float64(endProfit), float64(fromLeftProfits[i] + fromRightProfits[i])))
		if endProfit < fromLeftProfits[i] + fromRightProfits[i]{
			endProfit = fromLeftProfits[i] + fromRightProfits[i]
			fmt.Println(i)
		}
	}
	return endProfit
}