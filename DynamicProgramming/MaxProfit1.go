/*
@author : Aeishen
@data :  19/10/09, 21:29
@description :
*/

/*
题目：
    定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
注意你不能在买入股票前卖出股票。

示例 1:
	输入: [7,1,5,3,6,4]
	输出: 5
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
		 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:
	输入: [7,6,4,3,1]
	输出: 0
	解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

思路:这道题与IntOperation/MaxProfit2.go不同，这次只允许一次完成一笔交易。遍历所有时间的股票价格，以当前
	股票价格之前出现过的最低价作为买入价，并计算出当天价格出售的收益，与曾经的最大收益对比，遍历完即可的得到最大可能收益
*/

package main

import "fmt"

func main() {
	prices := []int{7,1,5,3,6,4}
	//prices := []int{1,2,3,4,5}
	//prices := []int{2,2,5}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices)> 0{
		minProfit := prices[0]
		for i:= 1;i < len(prices);i++{
			if prices[i] < minProfit{
				minProfit = prices[i]
			}

			// 当天出售收益：prices[i] - minProfit；曾经的最大收益：maxProfit
			if prices[i] - minProfit > maxProfit{
				maxProfit = prices[i] - minProfit
			}
		}
	}
	return maxProfit
}