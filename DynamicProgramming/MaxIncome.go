/*
@author : Aeishen
@data :  19/10/12, 16:28
@description :
*/

/*
题目：
	假设有8个任务需要完成，每个任务有各自的时间段，8个任务间任务段可能存在重合，如何选择任务使最后是收益最大
	下面为8个任务的时间段
	0   1   2   3   4   5   6   7   8   9   10   11
		******5******
				****1****
	************8************
					*******4*****
				***********6*********
						********3*******
							********2********
									*********4****
	*所覆盖的范围代表所经历的时间段，其中的数字代表这个任务可以带来的收益
*/

/*
思路：
	动态规划：
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	tast_incomes := []int{5,1,8,4,6,3,2,4}  // 每项任务的收入
	prev := []int{-1,-1,-1,0,-1,1,2,4}      // -1代表如果选择当前任务的话，之前没有可选的任务
    fmt.Println(maxIncome(tast_incomes,prev))
}

func maxIncome(incomes []int,prev []int) int{
	allResults := []int{incomes[0]}
	for i := 1;i < len(incomes); i++{
		if prev[i] == -1{
			if incomes[i] > allResults[i - 1] {
				allResults = append(allResults,incomes[i])
			}else{
				allResults = append(allResults,allResults[i - 1])
			}
		}else{
			if incomes[i] + allResults[prev[i]] > allResults[i - 1]{
				allResults = append(allResults,incomes[i] + allResults[prev[i]])
			}else{
				allResults = append(allResults,allResults[i - 1])
			}
		}
	}
	sort.Ints(allResults)
	return allResults[len(allResults) - 1]
}

func testMax(s []int)int{
	result := 0



	return result
}









