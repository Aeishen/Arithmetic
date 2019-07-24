/*
@author : Aeishen
@data :  19/07/24, 15:51

@description :
*/

/*
题目：
	求二叉树第k层的节点个数
*/

/*
解题思路：
     递归：
     非递归：
*/
package main

import (
	"Github/Arithmetic/DataStructure/Tree/binarySearchTree"
	"fmt"
)

//创建二叉树
func CreateTree_4() *Tree.MyBintree{
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
	arr := []int{5,15,3,7,13,17,1,8,12,19,16,6}

	for _,v := range arr{
		b.Add(v)
	}
	return b
}

//递归
func getKLevelNodeCount_1(root *Tree.Node, k int) int{
	if root == nil || k < 1{
		return 0
	}
	if k == 1{
		return 1
	}
	return getKLevelNodeCount_1(root.Left, k -1) + getKLevelNodeCount_1(root.Right, k -1)
}

//非递归
func getKLevelNodeCount_2(root *Tree.Node, k int) int {
	roots := []*Tree.Node{root}
	count := 1
	for count < k {
		deleteCount := len(roots)
		count++
		for i,_ := range roots{
			if roots[i].Right!=nil{
				roots = append(roots,roots[i].Right)
			}
			if roots[i].Left!=nil{
				roots = append(roots,roots[i].Left)
			}
		}
		if len(roots) == deleteCount{
			return 0
		}
		roots = roots[deleteCount:]
	}
	return len(roots)
}

func main() {
	b := CreateTree_4()
	fmt.Println(getKLevelNodeCount_2(b.Root, 5))
}
