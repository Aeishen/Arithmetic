/*
   @File : LongestUnivaluePath
   @Author: Aeishen
   @Date: 2019/12/26 21:38
   @Description:687. 最长同值路径
*/

/*
题目：
	给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:
	输入:
				  5
				 / \
				4   5
			   / \   \
			  1   1   5
	输出:2

示例 2:
	输入:
				  1
				 / \
				4   5
			   / \   \
			  4   4   5
	输出:2

思路：
     设计一个递归函数，返回以该节点为根节点，向下走的最长同值路径
     递归1：
        对于任意一个节点, 如果最长同值路径包含该节点, 那么只可能是两种情况:
        1. 其左右子树中加上该节点后所构成的同值路径中较长的那个继续向父节点回溯构成最长同值路径
        2. 左右子树加上该节点都在最长同值路径中, 构成了最终的最长同值路径
        需要注意因为要求同值, 所以在判断左右子树能构成的同值路径时要加入当前节点的值作为判断依据
			时间复杂度：O(N)，其中 N 是树中节点数。我们处理每个节点一次。
			空间复杂度：O(H)，其中 H 是树的高度。我们的递归调用栈可以达到 H 层的深度。
	  递归2:


*/

package main

import "fmt"

type TreeNode2 struct {
    Val int
    Left *TreeNode2
    Right *TreeNode2
}

func main() {
	//t1 := &TreeNode2{4,&TreeNode2{1,nil,nil},&TreeNode2{1,nil,nil}}
	//t2 := &TreeNode2{5,nil,&TreeNode2{5,nil,nil}}
	//root1 := &TreeNode2{5,t1,t2}
	//fmt.Println(longestUnivaluePath2(root1))
	//
	//t1 := &TreeNode2{4,&TreeNode2{4,nil,nil},&TreeNode2{4,nil,nil}}
	//t2 := &TreeNode2{5,nil,&TreeNode2{5,nil,nil}}
	//root1 := &TreeNode2{4,t1,t2}
	//fmt.Println(longestUnivaluePath2(root1))

	t1 := &TreeNode2{4,&TreeNode2{4,nil,nil},&TreeNode2{4,nil,nil}}
	t2 := &TreeNode2{4,nil,&TreeNode2{4,nil,nil}}
	root1 := &TreeNode2{4,t1,t2}
	fmt.Println(longestUnivaluePath2(root1))

	//root1 := &TreeNode2{1,nil,nil}
	//fmt.Println(longestUnivaluePath2(root1))
}

var ans = 0   // 记录全局最大值
func getMaxOne(a int,b int) int{
	if a > b{
		return a
	}
	return b
}


// 递归1
func longestUnivaluePath1(root *TreeNode2) int{
	ans = 0
	if root != nil {
		arrowLength(root,root.Val)
	}
	return ans
}

func arrowLength(root *TreeNode2,val int) int {
	if root == nil {
		return 0
	}

	l := arrowLength(root.Left,root.Val)
	r := arrowLength(root.Right,root.Val)
	ans = getMaxOne(ans, l + r)    // 路径长度为节点数减1所以此处不加1
	if root.Val == val{            // 和父节点值相同才返回以当前节点所能构成的最长同值路径长度, 否则返回0
		return getMaxOne(l,r) + 1
	}
	return 0
}

//---------------------------------------------------------------------------------

// 递归2
func longestUnivaluePath2(root *TreeNode2) int{
	ans = 0
	arrowLength1(root)
	return ans
}

func arrowLength1(root *TreeNode2) int {
	if root == nil {
		return 0
	}

	l := arrowLength1(root.Left)
	r := arrowLength1(root.Right)
	if root.Left != nil && root.Right != nil && root.Right.Val == root.Val && root.Left.Val == root.Val{
		ans = getMaxOne(ans, l + r + 2)
		return getMaxOne(l , r) + 1
	}

	if root.Left != nil  && root.Left.Val == root.Val{
		ans = getMaxOne(ans, l + 1)
		return l + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		ans = getMaxOne(ans, r + 1)
		return r + 1
	}

	return 0
}

// 递归3(递归2优化)
func longestUnivaluePath3(root *TreeNode2) int{
	ans = 0
	arrowLength2(root)
	return ans
}

//返回左右子树中相等节点距离最长的值
func arrowLength2(root *TreeNode2) int {
	if root == nil {
		return 0
	}

	l := arrowLength2(root.Left)
	r := arrowLength2(root.Right)
	pL,pR := 0,0
	if root.Left != nil  && root.Left.Val == root.Val{
		pL = l + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		pR = r + 1
	}
	ans = getMaxOne(ans, pL + pR)  //比较获取当前的相等的节点的最长长度
	return getMaxOne(pL, pR)       //返回左右子树中最大的那个子树
}