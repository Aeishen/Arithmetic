/*
   @File: ClimbStairs
   @Author: Aeishen
   @Date: 2019/10/09 20:34
   @Description: 爬楼梯
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
	动态规划：
		 不难发现，这个问题可以被分解为一些包含最优子结构的子问题，即它的最优解可以从其子问题的最优解来有效地构建，我们可以
		 使用动态规划来解决这一问题。第 i 阶可以由以下两种方法得到：
			在第 (i−1) 阶后向上爬 1 阶。
			在第 (i−2) 阶后向上爬 2 阶
		 所以到达第 i 阶的方法总数就是到第 (i-1)阶和第 (i−2) 阶的方法数之和。
		 令 dp[i] 表示能到达第 i 阶的方法总数：dp[i] = dp[i−1] + dp[i−2]

    递归：
		Arithmetic/Recursion/ClimbStairs.go

    斐波那契数：
		在上述方法中，我们使用 dpdp 数组，其中 dp[i]=dp[i-1]+dp[i-2]。可以很容易通过分析得出 dp[i] 其实就是第 i 个斐波那契数。
		Fib(n)=Fib(n-1)+Fib(n-2)

	斐波那契公式:
		直接使用斐波那契的数学公式，百度去，时间复杂度：O(log(n))。pow 方法将会用去 log(n) 的时间。空间复杂度：O(1)。使用常量级空间。

	Binets 方法：
		使用矩阵乘法来得到第 n 个斐波那契数，百度去，时间复杂度：O(log(n))。遍历 log(n) 位。空间复杂度：O(1)。使用常量级空间。
*/

package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(10))
	fmt.Println(climbStairs2(10))
}

// 动态规划，时间复杂度：O(n)，单循环到 n，空间复杂度：O(n)。dp 数组用了 n 的空间。
func climbStairs1(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int,n)
	dp[1] = 1
	dp[2] = 2
	if n <= 2{
		return dp[n]
	}else {
		for i := 3;i<=n;i++{
			dp[i] = dp[i - 1] + dp[i - 2]
		}
	}
	return dp[n]
}


// 斐波那契数，时间复杂度：O(n)。单循环到 n，需要计算第 n 个斐波那契数。空间复杂度：O(1)，使用常量级空间。
func climbStairs2(n int) int {
	if n <= 0 {
		return 0
	}

	curCount := 2
	preCount := 1
	if n <= 2{
		return n
	}else if n > 2{
		for i := 3;i<=n;i++{
			curCount = preCount + curCount
			preCount = curCount - preCount
		}
	}
	return curCount
}


