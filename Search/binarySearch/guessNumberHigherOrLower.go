/*
@author : Aeishen
@data :  19/07/18, 17:08
@description :
*/

/*
题目：
	We are playing the Guess Game. The game is as follows:
	I pick a number from 1 to n. You have to guess which number I picked.
	Every time you guess wrong, I'll tell you whether the number is higher or lower.

	You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):
	-1 : My number is lower
	 1 : My number is higher
	 0 : Congrats! You got it!

	Example:
	n = 10, I pick 6.
	Return 6.
*/

/*
解题思路：二分法
*/

package main

import "fmt"

var pickNum int

func guess(n int) int{
	compareResult := n - pickNum
	if compareResult < 0 {
		return -1
	}else if compareResult > 0 {
		return 1
	}else {
		return 0
	}
}

func guessNum(n int) int  {
	startNum := 1
	endNum := n

	for startNum <= endNum{
		mid := startNum + (endNum - startNum)/2
		fmt.Println(mid,startNum,endNum)
		guessResult := guess(mid)
		if guessResult == 1{
			endNum = mid - 1
		}else if guessResult == -1{
			startNum = mid + 1
		}else{
			return mid
		}
	}
	return -1
}

func main() {
	pickNum = 6
	fmt.Println(guessNum(10))
}
