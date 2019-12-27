/*
   @File: maxTreePathSum
   @Author: Aeishen
   @Date: 2019/12/27 15:29
   @Description: 124.二叉树中的最大路径和
*/

/*
题目：
	给定一个非空二叉树，返回其最大路径和。本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。
    该路径至少包含一个节点，且不一定经过根节点。


示例 1:
	输入: [1,2,3]
		   1
		  / \
		 2   3
	输出: 6

示例 2:
	输入: [-10,9,20,null,null,15,7]
		   -10
		   / \
		  9  20
			/  \
		   15   7
	输出: 42

思路：
    边界情况：如果节点为空，那么最大权值是 0 。
    递归1：
		二叉树 abc，a 是根结点（递归中的 root），bc 是左右子结点（代表其递归后的最优解）。
		最大的路径，可能的路径情况：
									a
								   / \
								  b   c
		1.b + a + c         ：表示如果不连接父结点，或本身是根结点的情况，这种情况是没法递归的，但是结果有可能是全局最大路径和。
		2.b + a + a 的父结点：
		3.a + c + a 的父结点：情况 2 和 3，递归时计算 a+b 和 a+c，选择一个更优的方案返回，也就是上面说的递归后的最优解。
		代码中使用 sum 来记录全局最大路径和。continuedSum 是情况 2 和 3。newSum 是情况 1。所要做的就是递归，递归时记录好全局最大和，返回到当前节点的一条最大路径。

	递归2：(解法与递归1一样，换种思路)
		对于每个节点，我们都会从两个方面思考
			作为新路径：该节点是该路径的root，此时的最大值来自于：node.left + node.right
			不是新路径：该节点不是路径的root，此时的最大值来自于：node.left 或 node.right
		所以对该节点的所有孩子递归调用 maxPath，计算左右子树各自最优解 l 和 r。
		检查是维护旧路径还是创建新路径。创建新路径的权值是：newPathSum = node.val + l + r，当新路径更好的时候更新 sum。
		而递归返回的到当前节点的一条最大路径，就是：node.val + max(l, r)。

	为什么递归2要比递归1少比较一种情况呢：
    	因为节点在取到最优路径时若为负数，则置0处理了，对于所有的 l 与 r 都有 l >= 0, r >= 0
        所以新路径的路径和：node.Val + l + r 永远 >=  node.val + max(l, r)
*/

package main

import (
	"fmt"
	"math"
)

type TreeNode3 struct {
	Val int
	Left *TreeNode3
	Right *TreeNode3
}


func main() {
	//t1 := &TreeNode3{2,nil,nil}
	//t2 := &TreeNode3{3,nil,nil}
	//root1 := &TreeNode3{1,t1,t2}
	//fmt.Println(maxPathSum(root1))

	t1 := &TreeNode3{9,nil,nil}
	t2 := &TreeNode3{20,&TreeNode3{15,nil,nil},&TreeNode3{7,nil,nil}}
	root1 := &TreeNode3{-10,t1,t2}
	fmt.Println(maxPathSum(root1))

	//root1 := &TreeNode3{1,nil,nil}
	//fmt.Println(maxPathSum(root1))
	//
	//t2 := &TreeNode3{-1,nil,nil}
	//root1 := &TreeNode3{2,nil,t2}
	//fmt.Println(maxPathSum(root1))
}

var sum int
func getMaxOne1(a int,b int) int{
	if a > b{
		return a
	}
	return b
}


// 递归1
func maxPathSum(root *TreeNode3) int {
	sum = math.MinInt32   // 结点值可能是负数，所以取最小值来比较
	if root != nil{
		maxPath(root)
	}
	return sum
}

func maxPath(root *TreeNode3)int  {
	if root == nil {
		return 0
	}
	l := getMaxOne1(maxPath(root.Left),0)
	r := getMaxOne1(maxPath(root.Right),0)
	newPathSum := root.Val + l + r
	sum = getMaxOne1(newPathSum,sum)
	return root.Val + getMaxOne1(l,r)
}

//递归2
func maxPathSum1(root *TreeNode3) int {
	sum = math.MinInt32   // 结点值可能是负数，所以取最小值来比较
	if root != nil{
		maxPath1(root)
	}
	return sum
}

func maxPath1(root *TreeNode3)int  {
	if root == nil {
		return 0
	}
	l := getMaxOne1(maxPath1(root.Left),0)   // 该结点下最优解如果是负数，则舍弃，不连接该路径，所以置0处理
	r := getMaxOne1(maxPath1(root.Right),0)  // 该结点下最优解如果是负数，则舍弃，不连接该路径，所以置0处理

	//此时的sum为原来所有路径的中的最大路径和
	newSum := root.Val + l + r                  // 新路径的路径和
	continuedSum := root.Val + getMaxOne1(l,r)  // 原来经过该节点路径的路径和接上该结点下最优路径和

	// 比较三种情况，获取并记录最大路径和
	sum = getMaxOne1(getMaxOne1(newSum,continuedSum),sum)
	return continuedSum
}