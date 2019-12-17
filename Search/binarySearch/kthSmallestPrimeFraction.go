/*
@author : Aeishen
@data :  19/11/08, 16:00
@description :
*/

/*
题目：
    一个已排序好的表 A，其包含 1 和其他一些素数.  当列表中的每一个 p<q 时，我们可以构造一个分数 p/q 。
    那么第 k 个最小的分数是多少呢?  以整数数组的形式返回你的答案, 这里 answer[0] = p 且 answer[1] = q.

示例:
	输入: A = [1, 2, 3, 5], K = 3
	输出: [2, 5]
	解释:
		已构造好的分数,排序后如下所示:
		1/5, 1/3, 2/5, 1/2, 3/5, 2/3.
		很明显第三个最小的分数是 2/5.

	输入: A = [1, 7], K = 1
	输出: [1, 7]

注意:
	A 的取值范围在 2 — 2000.
	每个 A[i] 的值在 1 —30000.
	K 取值范围为 1 —A.length * (A.length - 1) / 2
*/


package main

import "fmt"

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5},4))
	//fmt.Println(kthSmallestPrimeFraction([]int{1,5,7},2))
}

func kthSmallestPrimeFraction(A []int, K int) []int {
	left,right := float64(0),float64(1)
	mid := float64(0)
    p,q := 1,0
	count := 0
	for{
		mid = float64((left + right)/2)

		//A[i]为分子，A[j]为分母，寻找比mid小的分数个数
		count,p = 0,1
		for i := 0;i<len(A) - 1;i++ {
			isChange := false
			j := len(A) - 1
			fmt.Println("111 ",A[i],"/", A[j])
			for j > i && float64(A[i]) < mid * float64(A[j]){
				count ++
				fmt.Println("222 ",A[i],"/", A[j])
				j --
				isChange = true
			}


			//count = count + len(A) - j
			//if p * A[j] < q * A[i] {
			//	p = A[i]
			//	q = A[j]
			//}
            fmt.Println("-------------",A[i],"/", A[j],count,mid)
			if count == K {
				if isChange {
					j ++
				}
				return []int{A[i], A[j]}
			}
		}
        if count < K{
			left = mid
		}else{
			right = mid
		}
	}
	return []int{}
}