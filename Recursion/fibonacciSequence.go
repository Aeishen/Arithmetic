/*
@author : Aeishen
@data :  19/07/19, 10:23

@description : 斐波那契（Fibonacci）数列

问题背景：
	假定你有一雄一雌一对刚出生的兔子，它们在长到一个月大小时开始交配，在第二月结束时，雌兔子产下另一对兔子，过了一个月后它们也开始繁殖，如此这般持续下去。
	每只雌兔在开始繁殖时每月都产下一对兔子，假定没有兔子死亡，在一年后总共会有多少对兔子？

	在一月底，最初的一对兔子交配，但是还只有1对兔子；在二月底，雌兔产下一对兔子，共有2对兔子；在三月底，最老的雌兔产下第二对兔子，共有3对兔子；在四月底，
	最老的雌兔产下第三对兔子，两个月前生的雌兔产下一对兔子，共有5对兔子；……如此这般计算下去，兔子对数分别是：1,1,2,3,5,8,13,21,34,55,89,144,...
	看出规律了吗？从第3个数目开始，每个数目都是前面两个数目之和。这就是著名的斐波那契（Fibonacci）数列。
*/

package main

import (
	"fmt"
	"sync"
)

// 递归1: 逆向求解会造成大量的重复计算, 时间复杂度：O((3/2)^n)
func Fib_recursion_1(n int) int{
	if n <= 2{
		return 1    // 前两个
	}else{
		return Fib_recursion_1(n - 1) + Fib_recursion_1(n - 2)
	}
}

// 递归2：正向向求解可以得到近似线性的效率
func Fib_recursion_2(n int, preValue int, curValue int, curCount int) int{
	if curCount == n{
		return curValue  // 最后一个
	}else{
		curCount ++
		return Fib_recursion_2(n, curValue, preValue + curValue, curCount)
	}
}

// 迭代：用数组将每次计算的f(n)存储下来，用来下次计算用（空间换时间）,时间复杂度：O(n)
func Fib_iteration_3(n int)int{
	preValue := 0
	curValue := 1
	for i:= 1; i<n ;i++ {
		temp := curValue
		curValue = preValue + curValue
		preValue = temp
	}
	return curValue
}

func main() {
	var m sync.WaitGroup
	m.Add(3)
	go func() {
		fmt.Println("Fib_1:",Fib_recursion_1(20))
		m.Done()
	}()
	go func() {
		fmt.Println("Fib_3:",Fib_iteration_3(20))
		m.Done()
	}()
	go func() {
		fmt.Println("Fib_2:",Fib_recursion_2(20, 0, 1, 1))
		m.Done()
    }()
	m.Wait()
}