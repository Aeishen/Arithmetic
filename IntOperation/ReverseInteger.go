/*
@author : Aeishen
@data :  19/08/01, 14:46
@description :
*/

/*
题目：
     给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:
	输入: 123
	输出: 321

示例 2:
	输入: -123
	输出: -321

示例 3:
	输入: 120
	输出: 21

注意:
	假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) (value int) {
	for x != 0{
		value = value * 10 + x % 10
		x = x / 10
	}
	if value > math.MaxInt32 || value < math.MinInt32{
		value = 0
	}
	return
}

