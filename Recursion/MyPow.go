/*
   @File: MyPow
   @Author: Aeishen
   @Date: 2019/12/24 16:46
   @Description:
*/

/*
题目：
	实现 pow(x, n) ，即计算 x 的 n 次幂函数。

说明: -100.0 < x < 100.0, n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。

示例 1:
	输入: 2.00000, 10
	输出: 1024.00000

示例 2:
	输入: 2.10000, 3
	输出: 9.26100

示例 3:
	输入: 2.00000, -2
	输出: 0.25000
	解释: 2^-2 = (1/2)^2 = 1/4 = 0.25

思路：
    暴力：模拟将 x 相乘 n 次的过程。
		时间复杂度：O(n)。我们将 x 相乘 n 次
		空间复杂度：O(1)。我们需要一个变量来存储 x 的最终结果。

	递归快速幂算法：假设我们已经得到了 x ^ n的结果，那么如何才能算出 x ^ {2 * n }？显然，我们不需要再乘上 n 个 x。
                   使用公式 (x ^ n) ^ 2 = x ^ {2 * n}，只需要一次计算,我们就可以得到 x ^ {2 * n }。这种优化可以降低算法的时间复杂度。
		时间复杂度：O(log(n))。每次我们应用公式 (x ^ n) ^ 2 = x ^ {2 * n}，n 就减少一半。 因此，我们最多需要O(log(n)) 次计算来得到结果。
		空间复杂度：O(log(n))。每次计算，我们都需要存储 x ^ {n / 2}的结果。 我们需要计算 O(log(n)) 次，因此空间复杂度为 O(log(n))。

    迭代快速幂算法：使用公式 x ^ {a + b} = x ^ a * x ^ b
		时间复杂度：O(log(n))。对于 n 的每个二进制位，我们最多只能乘一次。所以总的时间复杂度为 O(log(n))。
		空间复杂度：O(1)。我们只需要两个变量来存储 x 的当前乘积和最终结果。
*/

package main

import "fmt"

func main() {
	fmt.Println(myPow2(2.00000, -2))
	fmt.Println(myPow2(2.00000, 10))
	fmt.Println(myPow2(2.10000, 3))
	fmt.Println(myPow2(2.10000, 0))
	fmt.Println(myPow2(2.0000, -1))
}

// 暴力
func myPow1(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans := float64(1)
	for i:= 0;i<n;i++{
		ans = ans * x
	}
	return ans
}

// 递归
func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64{
	if n == 0{
		return 1
	}
	half := fastPow(x, n / 2)
	if n % 2 == 0{
		return half * half
	}
	return half * half * x
}

// 迭代
func myPow3(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans := float64(1)
	currentProduct := x
	for  i := n; i > 0;{
		if (i % 2) == 1 {
			ans = ans * currentProduct
		}
		currentProduct = currentProduct * currentProduct
		i = i / 2
	}
	return ans
}
