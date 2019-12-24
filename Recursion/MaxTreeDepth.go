/*
   @File: MaxTreeDepth
   @Author: Aeishen
   @Date: 2019/12/24 15:44
   @Description:
*/

/*
题目：
	给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例:
	给定二叉树 [3,9,20,null,null,15,7]，
	3
   / \
  9  20
	/  \
   15   7
	返回它的最大深度3 。

思路：
	递归：每个节点的深度与它左右子树的深度有关，且等于其左右子树最大深度值加上 1, 可以写作：maxDepth(root) = max(maxDepth(root.left), maxDepth(root.right)) + 1
		时间复杂度：我们每个结点只访问一次，因此时间复杂度为 O(N)， 其中 N 是结点的数量。
		空间复杂度：在最糟糕的情况下，树是完全不平衡的，例如每个结点只剩下左子结点，递归将会被调用 N 次（树的高度），因此保持调用栈的存储将是 O(N)。
                   但在最好的情况下（树是完全平衡的），树的高度将是 log(N)。因此，在这种情况下的空间复杂度将是 O(log(N))。

    迭代：在栈的帮助下将上面的递归转换为迭代，使用 DFS 策略访问每个结点，同时在每次访问时更新最大深度。从包含根结点且相应深度为 1 的栈开始。
         然后继续迭代：将当前结点弹出栈并推入子结点。每一步都会更新深度。
		时间复杂度：O(N).
		空间复杂度：O(N).
*/

package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
	t := &TreeNode{3,&TreeNode{9,nil,nil},&TreeNode{20,&TreeNode{15,nil,nil},&TreeNode{7,nil,nil}}}
	//fmt.Println(maxDepth1(t))
	fmt.Println(maxDepth2(t))
}


// 递归
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxOne(maxDepth1(root.Left),maxDepth1(root.Right)) + 1
}

func maxOne(a int,b int) int{
	if a > b{
		return a
	}
	return b
}


type TreeNodeStack struct {
	depth int
	root *TreeNode
}

// 迭代
func maxDepth2(root *TreeNode) int {
	if root == nil{
		return 0
	}

	depth := 0
	stack := []TreeNodeStack{{1,root}}
	for len(stack) > 0{
		curDepth,curRoot := stack[len(stack) - 1].depth,stack[len(stack) - 1].root
		stack = stack[:len(stack) - 1]
		if curRoot != nil {
			depth = maxOne(depth,curDepth)
			stack = append(stack,TreeNodeStack{curDepth + 1,curRoot.Left})
			stack = append(stack,TreeNodeStack{curDepth + 1,curRoot.Right})
		}
	}
	
	return depth
}