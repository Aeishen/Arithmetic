/*
   @File : PascalTriangles1
   @Author: Aeishen
   @Date: 2019/12/16 22:36
   @Description: 杨辉三角 I
*/

/*
题目：
	给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
	在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
	输入: 5
	输出:
	[
		[1],
		[1,1],
		[1,2,1],
		[1,3,3,1],
		[1,4,6,4,1]
	]
*/
package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	all := make([][]int,0)
	return helper(numRows,all)
}


func helper(numRows int,all [][]int) [][]int {
	if numRows == 0 {
		return all
	}

	one := []int{1}
	l := len(all)
	if l > 0{
		one = make([]int,len(all[l - 1]) + 1)
		for i := 0;i < len(one);i++{
			if i == 0 || i == len(one) - 1{
				one[i] = 1
			}else{
				one[i] = all[l - 1][i - 1] + all[l - 1][i]
			}
		}
	}
	all = append(all,one)
	return helper(numRows - 1,all)
}
