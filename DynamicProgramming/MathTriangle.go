/*
@author : Aeishen
@data :  19/10/25, 10:35
@description :
*/

/*
题目：
    在给定的数字三角形中寻找一条从顶部到底边的路径，使得路径上所经过的数字之和最大。路径上的每一步都只能往左下或右下走。
    只需要求出这个最大和即可，不必给出具体路径。三角形的行数大于1小于等于100，数字为 0 - 99。

输入格式：
    5      //表示三角形的行数    接下来输入三角形
    7
    3   8
    8   1   0
    2   7   4   4
    4   5   2   6   5

思路:
     1.递归：用二维数组来存放数字三角形，然后用D(r,j)来表示第 r 行第 j 个数字(r,j从1开始算)，我们用MaxSum(r,j)表示从D(r,j)到底边的各条路径中，最佳路径的数字之和。
             因此，此题的最终问题就变成了求MaxSum(1,1)。从D(r, j)出发，下一步只能走D(r+1,j)或者D(r+1, j+1)。
     2.记忆递归型动态规划：在递归的基础上，将已经计算出的MaxSum(r,j)储存起来，下次用到时再直接取就可以了，免去重复计算，总的时间复杂度为n,因为三角形的数字总数为 n(n+1)/2。
     3.递推型动态规划：因为递归总是需要使用大量堆栈上的空间，很容易造成栈溢出，考虑把递归转换为递推
     4.优化递推：对于空间进行优化，没必要用二维curMaxSum数组存储每一个MaxSum(r,j),只要从底层一行行向上递推，那么只要一维数组maxSum[]即可,即只要存储一行的MaxSum值就可以
*/

package main

import (
	"fmt"
	"math"
)

var D = [][]int{
	{7},
	{3,8},
	{8,1,0},
	{2,7,4,4},
	{4,5,2,6,5},
}

var length int                       //三角形的高度，即行数
var curMaxSum map[int]map[int]int    //从第r行第j个数出发到底边路径的最大数字之和
var curLineMaxSum []int

func main() {
	length = len(D) - 1
	fmt.Println(recursion_MaxSum(0,0))

	curMaxSum = make(map[int]map[int]int)
	fmt.Println(memoryRecursion_MaxSum(0,0))

	curMaxSum = map[int]map[int]int{}
	fmt.Println(diTui_MaxSum())

	curLineMaxSum = make([]int,len(D[length]))
	fmt.Println(diTuiOptimize_MaxSum())
}

//递归
func recursion_MaxSum(r int, j int) (sum int) {
	if r == length {
		return D[r][j]
	}else{
		x :=  recursion_MaxSum(r + 1,j)       //表示左下
		y :=  recursion_MaxSum(r + 1,j + 1)   //表示右下
		return int(math.Max(float64(x),float64(y))) + D[r][j]
	}
}
/*
	拿第三行数字1来说，当我们计算从第2行的数字3开始的MaxSum时会计算出从1开始的MaxSum，当我们计算从第二行的数字8开始的MaxSum的时候又会计算一次从1开始的MaxSum，
    也就是说有重复计算。这样就浪费了大量的时间。也就是说如果采用递规的方法，深度遍历每条路径，存在大量重复计算。时间复杂度为2的n次方
*/


//记忆型递归
func memoryRecursion_MaxSum(r int, j int) (sum int) {
	if _,ok :=  curMaxSum[r][j];!ok{
		curMaxSum[r] = make(map[int]int)
		if r == length {
			curMaxSum[r][j] = D[r][j]
		}else{
			x :=  memoryRecursion_MaxSum(r + 1,j)       //表示左下
			y :=  memoryRecursion_MaxSum(r + 1,j + 1)   //表示右下
			curMaxSum[r][j] = int(math.Max(float64(x),float64(y))) + D[r][j]
		}
	}
	return curMaxSum[r][j]
}

// 递推
func diTui_MaxSum()(sum int){
	for r := length;r >= 0;r--{
		curMaxSum[r] = make(map[int]int)
		for j:= 0;j < len(D[r]);j++{
			if r == length{
				curMaxSum[length][j] = D[length][j]
			}else{
				curMaxSum[r][j] = D[r][j] + int(math.Max(float64(curMaxSum[r + 1][j]),float64(float64(curMaxSum[r + 1][j + 1]))))
			}
		}
	}
	return curMaxSum[0][0]
}


// 递推优化
func diTuiOptimize_MaxSum()(sum int){
	for r := length;r >= 0;r--{
		for j:= 0;j < len(D[r]);j++{
			if r == length{
				curLineMaxSum[j] = D[length][j]
			}else{
				curLineMaxSum[j] = D[r][j] + int(math.Max(float64(curLineMaxSum[j]),float64(float64(curLineMaxSum[j + 1]))))
			}
		}
	}
	return curMaxSum[0][0]
}