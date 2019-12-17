/*
   @File: ClimbStairs
   @Author: Aeishen
   @Date: 2019/12/17 15:54
   @Description:
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
    暴力递归：
	记忆化递归：
    正向递归：
	动态规划：  Arithmetic/DynamicProgramming/ClimbStairs.go
    斐波那契数：Arithmetic/DynamicProgramming/ClimbStairs.go
	斐波那契公式： Arithmetic/DynamicProgramming/ClimbStairs.go
    Binets 方法：  Arithmetic/DynamicProgramming/ClimbStairs.go
*/

package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(10))
	fmt.Println(climbStairs2(10))
	fmt.Println(climbStairs3(10,1,1,1))
}

// 暴力递归(会有大量重复计算)，时间复杂度：O(2^n)。树形递归的大小为2^n，空间复杂度：O(n)。递归树的深度可以达到 n
func climbStairs1(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 2{
		return n
	}
	return climbStairs1(n - 1) + climbStairs1(n - 2)
}


// 记忆化递归，用map缓存。时间复杂度：O(n) 。树形递归的大小可以达到 n，空间复杂度：O(n)。递归树的深度可以达到 n 。
var cacheMap = make(map[int]int)
func climbStairs2(n int) int {
	if n <= 0 {
		return 0
	}
	result, ok := cacheMap[n]
	if ok {
		return result
	}

	if n <= 2{
		result = n
	}else{
		result = climbStairs2(n - 1) + climbStairs2(n - 2)
	}
	cacheMap[n] = result
	return result
}


// 正向递归,正向向求解可以得到近似线性的效率
func climbStairs3(n int, pre int, cur int, count int) int {
	if n <= 0 {
		return 0
	}

	if n == count{
		return cur
	}else{
		count ++
		return climbStairs3(n, cur, pre + cur, count)
	}
}

