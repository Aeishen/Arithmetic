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

思路：递归：
		时间复杂度：O(N)，其中 N 是树中节点数。我们处理每个节点一次。
		空间复杂度：O(H)，其中 H 是树的高度。我们的递归调用栈可以达到 H 层的深度。


*/

package main

import "fmt"

type TreeNode2 struct {
    Val int
    Left *TreeNode2
    Right *TreeNode2
}

func main() {
	t1 := &TreeNode2{4,&TreeNode2{1,nil,nil},&TreeNode2{1,nil,nil}}
	t2 := &TreeNode2{5,nil,&TreeNode2{5,nil,nil}}
	root1 := &TreeNode2{5,t1,t2}
	fmt.Println(longestUnivaluePath(root1))

	t1 = &TreeNode2{4,&TreeNode2{4,nil,nil},&TreeNode2{4,nil,nil}}
	t2 = &TreeNode2{5,nil,&TreeNode2{5,nil,nil}}
	root1 = &TreeNode2{4,t1,t2}
	fmt.Println(longestUnivaluePath(root1))

	t1 = &TreeNode2{4,&TreeNode2{4,nil,nil},&TreeNode2{4,nil,nil}}
	t2 = &TreeNode2{5,nil,&TreeNode2{5,nil,nil}}
	root1 = &TreeNode2{1,t1,t2}
	fmt.Println(longestUnivaluePath(root1))
}

var ans = 0
func longestUnivaluePath(root *TreeNode2) int{
	arrowLength(root)
	return ans
}

func arrowLength(root *TreeNode2) int {
	if root == nil {
		return 0
	}

	l := arrowLength(root.Left)
	r := arrowLength(root.Right)
	arrowL, arrowR := 0, 0
	if root.Left != nil && root.Val == root.Left.Val{
		arrowL = arrowL + l + 1
	}
	if root.Right != nil && root.Val == root.Right.Val{
		arrowR = arrowR + r + 1
	}
	ans = getMaxOne(ans, arrowL + arrowR)
	return getMaxOne( arrowL , arrowR)
}

func getMaxOne(a int,b int) int{
	if a > b{
		return a
	}
	return b
}
