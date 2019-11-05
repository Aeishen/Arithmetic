/*
@author : Aeishen
@data :  19/10/12, 14:52
@description :斐波那契数列
*/

/*
思路：
    递归：fib(n) = fib(n - 1) + fib(n - 2) 使用递归的会产生大量的重复计算，比如要计算fib(5)，fib(5) = fib(4) + fib(3) 就要多次计算
          fib(3),fib(2)等之前计算过的值，浪费很多的空间和时间，当n值变得很大的时候，计算时间成指数增长，时间复杂度为O(2^n)。
	备忘录：在递归基础上将其之前计算过的fib值使用一个列表或者其他容器存储起来，当之后计算需要使用的时候就直接使用，不再去重复计算
           时间复杂度为O(2n+1) ---> O(n)
    动态规划：自底向上；把整个数想成一个数组，把a[0]+a[1]的值放在a[2]，然后a[1]+a[2]的值放在a[3]，以此类推，只要一个循环就可以得出答案了。
             整个方法因为没有递归的再调用，所以只被调用一次，从而大大减少了运行时间。时间复杂度为O(n)

区别：
    递归方法：采用的是自顶向下的思想，拆分为若干子问题，但是造成了子问题的重复求解。
    备忘录方法：采用的也是自顶向下的思想，但是该方法维护了一个记录子问题解的表，虽然填表动作的控制结构更像递归方法，但是的确避免了子问题的重复求解。
	动态规划思想：将原问题拆分为若干子问题，自底向上的求解。其总是充分利用重叠子问题，即通过每个子问题只解一次，把解保存在一个表中，巧妙的避免了子问题的重复求解。
*/

package main

import "fmt"

func main() {
	fmt.Println("fibonacci_Recursion1",fibonacci_Recursion1(20))
	fmt.Println("fibonacci_Recursion2",fibonacci_Recursion2(20,0,1,1))
	memo := make(map[int]int)
	fmt.Println("fibonacci_Memo",fibonacci_Memo(20, memo))
	fmt.Println("fibonacci_Dp",fibonacci_Dp(20))
}


// 递归1: 逆向求解会造成大量的重复计算, 时间复杂度：O((3/2)^n)
func fibonacci_Recursion1(n int) int{
	if n <= 2{
		return 1    // 前两个
	}else{
		return fibonacci_Recursion1(n - 1) + fibonacci_Recursion1(n - 2)
	}
}

// 递归2：正向向求解可以得到近似线性的效率
func fibonacci_Recursion2(n int, preValue int, curValue int, curCount int) int{
	if curCount == n{
		return curValue  // 最后一个
	}else{
		curCount ++
		return fibonacci_Recursion2(n, curValue, preValue + curValue, curCount)
	}
}

func fibonacci_Memo(n int, memo map[int]int) int {
	if v,isOk := memo[n];isOk{
		return v
	}
	result := 0
	if n > 0 {
		if n <= 2{
			result = 1
		}else{
			result =  fibonacci_Memo(n - 1,memo) + fibonacci_Memo(n - 2,memo)
		}
		memo[n] = result
	}
	return result
}

//
func fibonacci_Dp(n int) int {
	curResult := 1
	preResult := 0
	if n > 0 {
		for i := 1;i < n;i++{
			curResult = curResult + preResult
			preResult = curResult - preResult
		}
		return curResult
	}
	return 0
}
