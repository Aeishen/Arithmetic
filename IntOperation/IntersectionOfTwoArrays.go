/*
@author : Aeishen
@data :  19/09/04, 17:51
@description :
*/

/*
题目：
    给定两个数组，编写一个函数来计算它们的交集。

示例 1:
	输入: nums1 = [1,2,2,1], nums2 = [2,2]
	输出: [2,2]

示例 2:
	输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	输出: [4,9]

说明:
	输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。我们可以不考虑输出结果的顺序。

进阶:
	如果给定的数组已经排好序呢？你将如何优化你的算法？
	如果 nums1 的大小比 nums2 小很多，哪种方法更优？
	如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

思路1：用Map来建立nums1中字符和其出现个数之间的映射, 然后遍历nums2数组，如果当前字符在Map中的个数大于0，则将此字符加入结果res中，然后Map的对应值自减1。
思路2：给两个数组排序，然后用两个索引分别代表两个数组的起始位置，如果两个索引所代表的数字相等，则将数字存入结果中，两个索引均自增1，如果第一个索引所代表的数字大，则第二个索引自增1，反之亦然。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersect( []int{1,2,2,1}, []int{2,2}))
	fmt.Println(intersect( []int{4,9,5}, []int{9,4,9,8,4}))
	fmt.Println(intersect( []int{3,1,2}, []int{1,1}))

	fmt.Println(intersect1( []int{1,2,2,1}, []int{2,2}))
	fmt.Println(intersect1( []int{4,9,5}, []int{9,4,9,8,4}))
	fmt.Println(intersect1( []int{3,1,2}, []int{1,1}))
}

func intersect(nums1 []int, nums2 []int) []int {
	numMap := map[int]int{}
	same := []int{}
	for _,v := range nums1{
		numMap[v] ++
	}
	for _,v := range nums2{
		if numMap[v] > 0 {
			numMap[v] --
			same = append(same,v)
		}
	}
	return same
}

func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i,j := 0,0

	same := []int{}

	for i < len(nums1) && j <len(nums2){
		if nums1[i] > nums2[j]{
			j ++
		}else if nums1[i] < nums2[j]{
			i ++
		}else{
			same = append(same,nums1[i])
			i ++
			j ++
		}
	}
    return same
}
