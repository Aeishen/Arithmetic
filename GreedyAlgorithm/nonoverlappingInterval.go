/*
@author : Aeishen
@data :  19/07/26, 11:50
@description :
*/

/*
题目：
	给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠

注意:
    可以认为区间的终点总是大于它的起点。区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。



示例 1:
	输入: [ [1,2], [2,3], [3,4], [1,3] ]
	输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。




示例 2:
	输入: [ [1,2], [1,2], [1,2] ]
	输出: 2
解释:
	你需要移除两个 [1,2] 来使剩下的区间没有重叠。


示例 3:
	输入: [ [1,2], [2,3] ]
	输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
*/

/*
解题思路：动态规划：
        贪心算法：先排序按照顺序，两个指针遍历，一前一后，如果当前位置和上一个位置不冲突就顺序平移两个
                指针（后指针的值给前指针，然后后指针移动到下一位），如果冲突的话，那么前指针则变成当
				前两个指针当中覆盖最小的一个（贪心所在），后指针移动到下一个位置就好
*/

package main

import (
	"fmt"
	"sort"
)



// 贪心算法1:首先将输入的intervals按照end排序，然后保证我们每次放入区间的end最小，也就是对于后面要加入的区间留有更多的余地
type ints_1 [][]int

func (this ints_1)Less(i,j int)bool  {
	if this[i][len(this[i]) - 1] != this[j][len(this[j]) - 1]{
		return this[i][len(this[i]) - 1] < this[j][len(this[j]) - 1]
	}
	return this[i][0] < this[j][0]
}

func (this ints_1)Swap(i,j int)  {  this[i],this[j] = this[j],this[i]}

func (this ints_1)Len() int { return len(this) }

func countOfRemoveIntervals_1(intervals [][]int) int {
	l := len(intervals)
	if l == 0{
		return 0
	}
	s := ints_1(intervals)
	sort.Sort(s)                     // 排序按照每个区间最后一个元素大小升序排序
	countOfSaveInterval := 1         // 要保留的不重叠的区间的个数
	preValue := s[0][len(s[0]) - 1]  // 前一个区间的最后一个元素的值

	for i := 1;i < l; i++{
		if s[i][0] >= preValue { // 当前区间的第一个元素大于等于前一个区间的最后一个元素的值才能算不重叠
			preValue = s[i][len(s[i])-1]
			countOfSaveInterval ++
		}
	}
    return l - countOfSaveInterval
}


// 贪心算法2:将intervals按照start排序，然后遍历intervals，通过当前遍历的元素cur[0]和之前遍历的元素prev[1]比较，如果
//         prev[1] > cur[0]，那么显然cur区间不能放入，所以我们将result++，并且如果cur区间完全被prev包住，也就是
//         cur[1] < prev[1]，那么我们希望我们之前的区间中应该包括的是cur这个区间而不是prev，因为这样做我们可以为后面
//         的区间留有更多的余地。

type ints_2 [][]int

func (this ints_2)Less(i,j int)bool  {
	if this[i][0] != this[j][0]{
		return this[i][0] < this[j][0]
	}
	return this[i][len(this[i]) - 1] < this[j][len(this[j]) - 1]
}

func (this ints_2)Swap(i,j int)  {  this[i],this[j] = this[j],this[i]}

func (this ints_2)Len() int { return len(this) }

func countOfRemoveIntervals_2(intervals [][]int) int {
	l := len(intervals)
	if l == 0{
		return 0
	}
	s := ints_2(intervals)
	sort.Sort(s)                     // 排序按照每个区间最后一个元素大小升序排序
	countOfRomeInterval := 0         // 要删除的重叠的区间的个数
	preValue := s[0][len(s[0]) - 1]  // 前一个区间的最后一个元素的值

	for i := 1;i < l; i++{
		if s[i][0] < preValue {
			countOfRomeInterval ++
			if s[i][len(s[i]) - 1] <= preValue{
				preValue = s[i][len(s[i]) - 1]
			}
		}else{
			preValue = s[i][len(s[i]) - 1]
		}
	}
	return countOfRomeInterval
}


// 动态规划-------------------

func countOfRemoveIntervals_3(intervals [][]int) int {
	l := len(intervals)
	if l == 0{
		return 0
	}
	s := ints_2(intervals)
	sort.Sort(s)                     // 排序按照每个区间最后一个元素大小升序排序
	countOfRomeInterval := 0         // 要删除的重叠的区间的个数
	preValue := s[0][len(s[0]) - 1]  // 前一个区间的最后一个元素的值

	for i := 1;i < l; i++{
		if s[i][0] < preValue {
			countOfRomeInterval ++
			if s[i][len(s[i]) - 1] < preValue{
				preValue = s[i][len(s[i]) - 1]
			}
		}else{
			preValue = s[i][len(s[i]) - 1]
		}
	}
	return countOfRomeInterval
}

func main() {
	//intervals := [][]int{{1,2},{2,3},{3,4},{1,3}}
	intervals := [][]int{{1,2},{1,2},{1,2},{2,3},{2,5},{3,4}}
	//intervals := [][]int{{1,2},{2,3}}
	fmt.Println(countOfRemoveIntervals_1(intervals))
	fmt.Println(countOfRemoveIntervals_2(intervals))
}