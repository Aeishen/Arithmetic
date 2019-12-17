/*
   @File : PascalTriangles2
   @Author: Aeishen
   @Date: 2019/12/16 23:17
   @Description: 杨辉三角 II
*/

/*
题目：
	给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
	在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
	输入: 3
	输出: [1,2,1]

进阶：
	你可以优化你的算法到 O(k) 空间复杂度吗？
*/

package main

import "fmt"

func main() {
	fmt.Println(getRow(2))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0{
		return []int{1}
	}
	return helper3(rowIndex,[]int{1})
}

func helper2(rowIndex int,one []int)[]int{
	if rowIndex == 0{
		return one
	}
	one = append(one,1)
	anotherOne := make([]int,len(one))
	copy(anotherOne,one)
	for i:= 1 ;i < len(one) - 1;i++{
		anotherOne[i] = one[i - 1] + one[i]
	}
	return helper2(rowIndex - 1,anotherOne)
}

func helper3(rowIndex int,one []int)[]int{
	if rowIndex == 0{
		return one
	}
	one = append(one,1)
	for i:= len(one) - 2 ;i > 0;i--{
		one[i] = one[i - 1] + one[i]
	}
	fmt.Println()
	return helper3(rowIndex - 1,one)
}