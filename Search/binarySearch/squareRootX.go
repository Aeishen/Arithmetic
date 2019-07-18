/*
@author : Aeishen
@data :  19/07/18, 17:29
@description :Implement int sqrt(int x)
*/

/*
题目：
    Compute and return the square root of x.
*/

/*
解题思路：二分法
*/

package main

import "fmt"

func sqrt(x int) int{
	startNum := 1
	endNum := x

	for startNum <= endNum{
		mid := startNum + (endNum - startNum)/2
		if mid == x / mid && x % mid == 0{
			return mid
		}else if x > mid * mid {
			startNum = mid + 1
		}else if x < mid * mid {
			endNum = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(sqrt(66))
	fmt.Println(sqrt(64))
	fmt.Println(sqrt(33))
	fmt.Println(sqrt(4))
	fmt.Println(sqrt(8))
	fmt.Println(sqrt(25))
}
