/*
@author : Aeishen
@data :  19/09/04, 16:40
@description :
*/

/*
题目：
    给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：
	你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
	输入: [7,1,5,3,6,4]
	输出: 7
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     	随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

示例 2:
	输入: [1,2,3,4,5]
	输出: 4
	解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
		 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
		 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

示例 3:
	输入: [7,6,4,3,1]
	输出: 0
	解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


思路1：峰谷法
思路2：贪心：我们要算的是利润，要有利润，自然要有一次交易。按照贪心策略，不关心以后，我们只关心当前利益。
             第0天买入，花费prices[0]，第一天卖出，得到prices[1]，那么我们的收获就是curMoney = prices[1] - prices[0],那么有两种情况
			（1）当profit > 0 时，赶紧买入卖出，能赚一笔是一笔，苍蝇再小也是肉嘛 
			（2）当profit <= 0 时，再买入卖出的话，那就是傻了，白费力气不说，还亏钱。
			以此方式类推下去，即得最大利润。该思路与最大子序列和很是相像，可以试试独立快速完成查看是否理解

*/
package main

import "fmt"

func main() {
	prices := []int{7,1,5,3,6,4}
	//prices := []int{1,2,3,4,5}
	//prices := []int{2,2,5}
	fmt.Println(maxProfit2(prices))
}


func maxProfit1(prices []int) int {
	buyIndex,saleIndex := -1,-1
	money := 0
	if len(prices) > 1{
		for i := 0;i < len(prices); {
			if i == 0 {
				if prices[i] < prices[i+1]{
					buyIndex = i
				}
			}else if i < len(prices) - 1{
				if buyIndex == -1 && prices[i] <= prices[i+1] && prices[i] <= prices[i-1]{
					buyIndex = i
				}else if buyIndex != -1 && prices[i] >= prices[i+1] && prices[i] >= prices[i-1]{
					saleIndex = i
				}
			}else{
				if buyIndex != -1 && prices[i] > prices[i-1]{
					saleIndex = i
				}
			}
			if saleIndex != -1 && buyIndex != -1{
				money = money + prices[saleIndex] - prices[buyIndex]
				saleIndex = -1
				buyIndex = -1
			}
		}
	}

	return money
}

func maxProfit2(prices []int) int {
	money := 0
	if len(prices) > 1{
		for i := 1;i < len(prices);i ++ {
			curMoney := prices[i] - prices[i - 1]
			if curMoney > 0 {
				money = money + curMoney
			}
		}
	}
	return money
}