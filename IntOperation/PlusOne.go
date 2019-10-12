/*
@author : Aeishen
@data :  19/09/06, 09:44
@description : leetcode 66
*/

/*
题目：
    给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:
	输入: [1,2,3]
	输出: [1,2,4]
	解释: 输入数组表示数字 123。

示例 2:
	输入: [4,3,2,1]
	输出: [4,3,2,2]
	解释: 输入数组表示数字 4321。

思路：分情况，加1后小于10；加1后为10
*/


package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,9}))
}

func plusOne(digits []int) []int {
    for i := len(digits) - 1;i >= 0; i--{
		digits[i] ++
		if digits[i] % 10 == 0{
			if i == 0 {
				digits = append(digits,0)
				digits[0] = 1
			}else{
				digits[i] = 0
			}
		}else{
			break
		}
	}
    return digits
}