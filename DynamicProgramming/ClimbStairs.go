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
	动态规划：因为每次只能爬 1 或 2 个台阶。假设知道爬上 n - 1 阶的方法为 counts_1, 则从 n - 1 阶爬上 n 阶也是 counts_1 种方法
            （每种方法都是爬 1 个台阶就可以到 n 阶）；假设知道爬上 n - 2 阶的方法为 counts_2, 则从 n - 2 阶爬上 n 阶也是 counts_2 种方法。
			（每种方法都是爬 2 个台阶就可以到 n 阶）；所以爬上 n 阶总的方法有 counts_1 + counts_2 种
*/

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	curCount := 0      // 爬上 n 阶的方法数
	preCount := 2      // 爬上 n - 1 阶的方法数
	prePreCount := 1   // 爬上 n - 2 阶的方法数
	if n == 1 {
		return prePreCount
	}else if n == 2{
		return preCount
	}else if n > 2{
		for i := 3;i<=n;i++{
			curCount = preCount + prePreCount
			prePreCount = preCount
			preCount = curCount
		}
	}
	return curCount
}