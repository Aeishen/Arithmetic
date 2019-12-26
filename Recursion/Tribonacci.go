/*
   @File : Tribonacci
   @Author: Aeishen
   @Date: 2019/12/25 22:08
   @Description: 1137. 第 N 个泰波那契数
*/

/*
题目：
	泰波那契序列 Tn 定义如下：
	T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
	给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

示例 1:
	输入：n = 4
	输出：4
	解释：
	T_3 = 0 + 1 + 1 = 2
	T_4 = 1 + 1 + 2 = 4

示例 2:
	输入：n = 25
	输出：1389537

提示：
	0 <= n <= 37
	答案保证是一个 32 位整数，即 answer <= 2^31 - 1。

思路：参考 fibonacciSequence.go

*/

package main

import "fmt"

func main() {
	fmt.Println(tribonacci1(4))
	fmt.Println(tribonacci1(25))
	fmt.Println(tribonacci2(4))
	fmt.Println(tribonacci2(25))
}

// 记忆化递归
func tribonacci1(n int) int {
	m := make(map[int]int)
	return helper4(n,m)
}

func helper4(n int,m map[int]int) int{
	if n <= 0{
		return 0
	}else if n <= 2{
		return 1
	}
	if v,ok := m[n];ok{
		return v
	}

	m[n] = helper4(n - 1,m) + helper4(n - 2,m) + helper4(n - 3,m)
	return m[n]
}

// 迭代
func tribonacci2(n int) int {
	if n <= 0{
		return 0
	}else if n <= 2{
		return 1
	}

	curVal := 1
	preVal := 1
	prePreVal := 0
	for i := 3;i <= n;i++{
		curVal = curVal + preVal + prePreVal
		preVal = curVal - prePreVal - preVal
		prePreVal = curVal - prePreVal - preVal
	}
	return curVal
}
