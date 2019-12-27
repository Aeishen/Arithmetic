/*
   @File: DiameterOfBinaryTree
   @Author: Aeishen
   @Date: 2019/12/27 17:36
   @Description: 543.二叉树的直径
*/

/*
题目：
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。


示例 ：
	给定二叉树

			  1
			 / \
			2   3
		   / \
		  4   5
	返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。

思路：
	递归：任意一条路径可以被写成两个箭头（不同方向），每个箭头代表一条从某些点向下遍历到孩子节点的路径。
	假设我们知道对于每个节点最长箭头距离分别为 L, R，那么最优路径经过 L + R + 1 个节点。
    按照常用方法计算一个节点的深度：max(depth of node.left, depth of node.right) + 1。在计算的同时，
    经过这个节点的路径长度为 1 + (depth of node.left) + (depth of node.right) 。
    搜索每个节点并记录这些路径经过的点数最大值，期望长度是获得的节点数 - 1。
*/
package main

import "fmt"

type TreeNode4 struct {
	Val int
	Left *TreeNode4
	Right *TreeNode4
}



func main() {
	//t1 := &TreeNode4{3,nil,nil}
	//t2 := &TreeNode4{2,&TreeNode4{4,nil,nil},&TreeNode4{5,nil,nil}}
	//root1 := &TreeNode4{1,t1,t2}
	//fmt.Println(diameterOfBinaryTree(root1))

	t1 := &TreeNode4{3,nil,nil}
	t2 := &TreeNode4{2,nil,nil}
	root1 := &TreeNode4{1,t1,t2}
	fmt.Println(diameterOfBinaryTree(root1))
}
var lenPath int
func getMaxOne2(a int,b int) int{
	if a > b{
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode4) int {
	lenPath = 1
	if root != nil{
		diameter(root)
	}
	return lenPath - 1
}

func diameter(root *TreeNode4) int{
	if root == nil{
		return 0
	}

	l := diameter(root.Left)   // 左子树的最多节点路径的节点数
	r := diameter(root.Right)  // 右子树的最多节点路径的节点数
	lenPath = getMaxOne2(lenPath,l + r + 1)

	return getMaxOne2(l , r) + 1
}