/*
@author : Aeishen
@data :  19/07/23, 11:17

@description :
*/

/*
题目：
	翻转一颗二叉树（即所有左右节点互换）
*/

/*
解题思路：
     递归：交换当前左右节点，递归
*/

package main

import (
	"./binarySearchTree"
	"Github/Arithmetic/DataStructure/Queue/queue"
	"fmt"
)

//创建二叉树
func createTree() *Tree.MyBintree{
	b := Tree.NewMyBintree(10)
	b.CompareBig = func(data1 Tree.Object, data2 Tree.Object) int {
		if data1.(int) > data2.(int) {
			return 1
		}else if data1.(int) < data2.(int) {
			return -1
		}else{
			return 0
		}
	}
	arr := []int{5,15,3,7,13,17}

	for _,v := range arr{
		b.Add(v)
	}
	return b
}

//翻转二叉树
func reversalTree(b *Tree.MyBintree) {
	//b.Root = reversal_rec(b.Root)
	reversal_unrec(b.Root)
}

// 递归翻转
func reversal_rec(root *Tree.Node) *Tree.Node{
	if root.Right == nil && root.Left == nil {
		return root
	}
	root.Left ,root.Right = root.Right,root.Left
	root.Left = reversal_rec(root.Left)
	root.Right = reversal_rec(root.Right)
	return root
}

//非递归翻转:跟二叉树的层序遍历一样，需要用 queue 来辅助。
//首先把根节点排入队列中，然后从队中取出来，交换其左右节点，左右节点如果存在则分别将左右节点在排入队列中，
//以此类推直到队列中没有节点了再停止循环，即可。
func reversal_unrec(root *Tree.Node) {
   q := Queue.NewQueue()
   q.Enqueue(root)
   for !q.IsEmoty(){
   	   curRoot := q.GetFront().(*Tree.Node)
   	   q.Dequeue()
	   temp := curRoot.Left
	   curRoot.Left = curRoot.Right
	   curRoot.Right = temp

	   if curRoot.Left != nil{
	   		q.Enqueue(curRoot.Left)
	   }
	   if curRoot.Right != nil{
		   q.Enqueue(curRoot.Right)
	   }
   }
}

func main() {
	b := createTree()
	fmt.Println(b.InorderTraversal())
	reversalTree(b)
	fmt.Println(b.InorderTraversal())

}