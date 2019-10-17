/*
@author : Aeishen
@data :  19/07/26, 10:44
@description :
*/

/*
题目：
	假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 gi ，
    这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分
    配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

注意:
    你可以假设胃口值为正。一个小朋友最多只能拥有一块饼干。



示例 1:
	输入: [1,2,3], [1,1]
	输出: 1

解释:
	你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
	虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
	所以你应该输出1。


示例 2:
	输入: [1,2], [1,2,3]
	输出: 2

解释:
	你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
	你拥有的饼干数量和尺寸都足以让所有孩子满足。
	所以你应该输出2.
*/

/*
解题思路：贪心算法，把最大的饼干分给能满足的最贪心的朋友
*/
package main

import (
	"fmt"
	"sort"
)

func maximumSatisfaction(children []int,biscuits []int)(maxCount int){

	sort.Ints(children)
	sort.Ints(biscuits)

	len_child := len(children)
	len_biscuit := len(biscuits)

	len_biscuit--
	len_child --
	for len_biscuit >= 0 && len_child >= 0{
		if biscuits[len_biscuit] >= children[len_child]{
			maxCount ++
			len_biscuit --
		}
		len_child --   //找下一个孩子的为胃口值
	}
	return
}

func main() {
	//childrens := []int{1,2}
	//biscuits := []int{1,2,3}
	childrens := []int{1,2,3}
	biscuits := []int{1,1}
	fmt.Println(maximumSatisfaction(childrens,biscuits))
}
