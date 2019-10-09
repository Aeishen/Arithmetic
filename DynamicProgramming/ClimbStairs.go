/*
@author : Aeishen
@data :  19/10/09, 20:34
@description :
*/

/*
题目：
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：
	输入： 2
	输出： 2
	解释： 有两种方法可以爬到楼顶。
	1.  1 阶 + 1 阶
	2.  2 阶

示例 2：
	输入： 3
	输出： 3
	解释： 有三种方法可以爬到楼顶。
	1.  1 阶 + 1 阶 + 1 阶
	2.  1 阶 + 2 阶
	3.  2 阶 + 1 阶
*/

/*
思路：
	动态规划：爬上第n阶之前，只有两种可能，爬了一阶或者两阶，因此此时他总共可拥有的次数，即为他最后爬了一阶情况下原来n-1阶的次数加上最后爬了两阶时
    原来n-2阶的次数。
*/

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	count := 0
	preCount := 2
	prePreCount := 1
	if n == 1 {
		return prePreCount
	}else if n == 2{
		return preCount
	}else if n > 2{
		for i := 3;i<=n;i++{
			count = preCount + prePreCount
			prePreCount = preCount
			preCount = count
		}
	}
	return count
}