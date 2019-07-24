/*
@author : Aeishen
@data :  19/07/24, 10:37

@description :
*/

/*
题目：
	获取二叉树的深度
*/

/*
解题思路：
     递归：对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
     非递归：借助队列，在进行按层遍历时，记录遍历的层数即可
*/

package main

import (
	"Github/Arithmetic/DataStructure/Tree/binarySearchTree"
	"fmt"
	"sync"
)
//创建二叉树
func CreateTree_1() *Tree.MyBintree{
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

//递归实现1
func GetTreeDepth_rec_1(root *Tree.Node) int{
	if root == nil {
		return 0
	}

	countL := GetTreeDepth_rec_1(root.Left)
	countR := GetTreeDepth_rec_1(root.Left)

	if countL > countR{
		return countL + 1
	}
	return countR + 1
}

//递归实现2
func GetTreeDepth_rec_2(root *Tree.Node) int{
	if root == nil {
		return 0
	}

	countL := GetTreeDepth_rec_2(root.Left) + 1
	countR := GetTreeDepth_rec_2(root.Left) + 1

	if countL > countR{
		return countL
	}
	return countR
}

// 非递归实现
func GetTreeDepth_unrec(root *Tree.Node) int{
	if root == nil {
		return 0
	}
	depth := 0
	nodes := []*Tree.Node{root}
	for len(nodes) >0 {
		depth ++
		size := len(nodes) // 每层节点数
		count := 0
		for count < size{
			curNode := nodes[0]
			nodes = nodes[1:]
			count ++

			if curNode.Left != nil{
				nodes = append(nodes,curNode.Left)
			}
			if curNode.Right != nil{
				nodes = append(nodes,curNode.Right)
			}
		}
	}
	return depth
}

func main() {
	var w sync.WaitGroup
	w.Add(3)
	b := CreateTree_1()
	go func() {
		fmt.Println("递归1：",GetTreeDepth_rec_1(b.Root))
		w.Done()
	}()
	go func() {
		fmt.Println("递归2：",GetTreeDepth_rec_2(b.Root))
		w.Done()
	}()
	go func() {
		fmt.Println("非递归：",GetTreeDepth_unrec(b.Root))
		w.Done()
	}()
	w.Wait()

}