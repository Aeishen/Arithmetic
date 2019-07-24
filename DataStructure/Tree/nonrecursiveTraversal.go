/*
@author : Aeishen
@data :  19/07/24, 11:27

@description :
*/

/*
题目：
	使用非递归算法实现树的前中后序三种遍历与层次遍历
*/

/*
解题思路：
     前中后序三种遍历：借助栈的先进后出实现
	 层次遍历：借助队列先进先出实现
*/

package main

import (
	"Github/Arithmetic/DataStructure/Queue/queue"
	"Github/Arithmetic/DataStructure/Stack/stack"
	"Github/Arithmetic/DataStructure/Tree/binarySearchTree"
	"fmt"
)

//创建二叉树
func CreateTree_2() *Tree.MyBintree{
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

// 前序遍历
func preorderTraversal(root *Tree.Node) (all []Stack.Item) {
	s := Stack.NewStack()
	curRoot := root
	for !s.IsEmpty() || curRoot != nil{
		for curRoot != nil{      // 从根结点开始，寻找左子树，把它压入栈中
			all = append(all, curRoot.Data)
			s.Push(curRoot)
			curRoot = curRoot.Left
		}
		curRoot = s.Peek().(*Tree.Node)
		s.Pop()
		curRoot = curRoot.Right  // 然后开始寻找右子树
	}
	return
}


// 中序遍历
func inorderTraversal(root *Tree.Node) (all []Stack.Item) {
	s := Stack.NewStack()
	curRoot := root
	for !s.IsEmpty() || curRoot != nil {
		for curRoot != nil{      // 从根结点开始，寻找左子树，把它压入栈中
			s.Push(curRoot)
			curRoot = curRoot.Left
		}
		curRoot = s.Peek().(*Tree.Node)
		all = append(all, curRoot.Data)
		s.Pop()
		curRoot = curRoot.Right  // 然后开始寻找右子树
	}
	return
}

// 后序遍历
func postorderTraversal(root *Tree.Node) (all []Stack.Item) {
	s1 := Stack.NewStack()
	s2 := Stack.NewStack()
	s1.Push(root)

	for !s1.IsEmpty(){
		curRoot := s1.Peek().(*Tree.Node)
		s1.Pop()
		if curRoot.Left != nil{
			s1.Push(curRoot.Left)
		}
		if curRoot.Right != nil{
			s1.Push(curRoot.Right)
		}
		s2.Push(curRoot)
	}
	for !s2.IsEmpty(){
		all = append(all,s2.Peek().(*Tree.Node).Data)
		s2.Pop()
	}
	return
}

// 层次遍历1
func levelTraversal_1(root *Tree.Node) (all []Stack.Item) {
	q := Queue.NewQueue()
	q.Enqueue(root)
	roots := []*Tree.Node{}
	for !q.IsEmoty(){
		curRoot := q.GetFront().(*Tree.Node)
		all = append(all,curRoot.Data)
		roots = append(roots,curRoot)
		q.Dequeue()
		if q.IsEmoty(){
			for i,_ := range roots{
				if roots[i].Left != nil{
					q.Enqueue(roots[i].Left)
				}
				if roots[i].Right != nil{
					q.Enqueue(roots[i].Right)
				}
			}
			roots = []*Tree.Node{}
		}
	}
	return
}

// 层次遍历2
func levelTraversal_2(root *Tree.Node) (all []Stack.Item) {
	roots := []*Tree.Node{root}
	for len(roots)>0{
		curRoot := roots[0]
		roots = roots[1:]
		all = append(all,curRoot.Data)

		if curRoot.Left != nil{
			roots = append(roots,curRoot.Left)
		}
		if curRoot.Right != nil{
			roots = append(roots,curRoot.Right)
		}
	}
	return
}

func main() {
	b := CreateTree_2()
	fmt.Println(preorderTraversal(b.Root))
	fmt.Println(inorderTraversal(b.Root))
	fmt.Println(postorderTraversal(b.Root))
	fmt.Println(levelTraversal_1(b.Root))
	fmt.Println(levelTraversal_2(b.Root))
}