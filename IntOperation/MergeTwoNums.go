/*
   @File : MergeTwoNums
   @Author: Aeishen
   @Date: 2020/1/7 23:26
   @Description: 88. 合并两个有序数组
*/

/*
题目：
	给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。


示例 ：
	输入:
		nums1 = [1,2,3,0,0,0], m = 3
		nums2 = [2,5,6],       n = 3
    输出:
		[1,2,2,3,5,6]

说明:
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

思路：
	nums1 长度足够容纳 nums1 现有的元素和 nums2 现有的元素，那么合并后 nums1 最后一个元素一定是 nums1 和 nums2 中
    最大元素。依次类推，由于 nums1 和 nums2 都已经是升序直接将，将 nums1 和 nums2 的元素从后往前逐个比较取大的放入
    nums1 中对应位置即可，当 nums1 和 nums2 其中一个原来的元素已经取完，则将剩余一个的元素直接填充在合并后 nums1
    对应位置即可
*/
package main

import "fmt"

func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1 , m , nums2 , n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums1) > m + n{
		return
	}
    i ,j := m -1, n -1
    last := m + n -1  // 标记要填充最大值的位置
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j]{
			nums1[last] = nums2[j]
			j --
		}else{
			nums1[last] = nums1[i]
			i --
		}
		last --
	}

	for k := last; k >= 0 && j >= 0;k --{
		nums1[k] = nums2[j]
		j --
	}

}