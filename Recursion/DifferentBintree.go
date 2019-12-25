/*
   @File: DifferentBintree
   @Author: Aeishen
   @Date: 2019/12/24 17:52
   @Description:
*/

/*
题目：
	给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:
	输入: 3
	输出:
	[
	  [1,null,3,2],
	  [3,2,null,1],
	  [3,1,null,null,2],
	  [2,1,3],
	  [1,null,2,null,3]
	]
	解释:
	以上的输出对应以下 5 种不同结构的二叉搜索树：
	   1         3     3      2      1
		\       /     /      / \      \
		 3     2     1      1   3      2
		/     /       \                 \
	   2     1         2                 3

思路：
	递归：从序列 1 ..n 取出数字 i 并以它作为当前树的根节点。 那么就有 i - 1 个元素可以用来构造左子树，而另外的 n - i 个元素可
		以用于构造右子树。根据左右子树的返回的所有子树，依次选取然后接上（每个左边的子树跟所有右边的子树匹配，而每个右边的子树
       也要跟所有的左边子树匹配，总共有左右子树数量的乘积种情况），构造好之后作为当前树的结果返回。
*/
package main

import "fmt"

type TreeNode1 struct {
	Val int
	Left *TreeNode1
	Right *TreeNode1
}

func main() {
	fmt.Println(generateTrees(3))
	fmt.Println(MemGenerateTrees(3))
}

// 正常递归
func generateTrees(n int) []*TreeNode1 {
	if n == 0 {
		return nil
	}
	return generate_trees(1 ,n)
}

func generate_trees(sta int ,en int)[]*TreeNode1{
	if sta > en{
		return []*TreeNode1{nil}
	}
	var allTrees []*TreeNode1
	for i := sta; i <= en; i ++{
		leftTrees := generate_trees(sta ,i - 1)
		rightTrees := generate_trees(i + 1,en)
		for _,l := range leftTrees{
			for _,r := range rightTrees{
				curTree := &TreeNode1{i,l,r}
				allTrees = append(allTrees,curTree)
			}
		}
	}
	return allTrees
}


// 记忆化递归
func MemGenerateTrees(n int) []*TreeNode1 {
	if n <= 0 {
		return nil
	}
	mem := make(map[[2]int][]*TreeNode1)
	mem_generate_trees(1,n,mem)
	return mem[[2]int{1,n}]
}

func mem_generate_trees(sta int ,en int,mem map[[2]int][]*TreeNode1)[]*TreeNode1{
	if sta > en {
		return []*TreeNode1{nil}
	}

	if v,ok := mem[[2]int{sta,en}];ok{
		return v
	}

	var allTrees []*TreeNode1

	for i:= sta;i<= en;i++{
		leftTrees := mem_generate_trees(sta,i - 1,mem)
		rightTrees := mem_generate_trees(i + 1,en,mem)
		for _,l := range leftTrees{
			for _,r := range rightTrees{
				curT := &TreeNode1{i,l,r}
				allTrees = append(allTrees,curT)
			}
		}
	}
	mem[[2]int{sta,en}] = allTrees
	return allTrees
}
