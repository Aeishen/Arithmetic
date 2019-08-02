/*
@author : Aeishen
@data :  19/07/18, 15:20

@description :
*/

/*
题目：
	Given a non-empty binary search tree and a target value, find the value in the BST that is closest to the target.

	Note:
		 Given target value is a floating point.
         You are guaranteed to have only one unique value in the BST that is closest to the target.
*/

/*
解题思路：
     找到最接近这个值的两个节点（父子节点），再比较哪个更加接近
*/

package main

import (
	"Github/Arithmetic/DataStructure/Tree/binarySearchTree"
	"fmt"
)

//创建二叉树
func CreateTree_5(arr []int) *Tree.MyBintree{
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


	for _,v := range arr{
		b.Add(v)
	}
	return b
}

func checkIsSame(root1 *Tree.Node,root2 *Tree.Node)bool{
	if root1 == nil && root2 == nil {
		return true
	}else if  root1 == nil || root2 == nil{
		return false
	}
	return checkIsSame(root1.Left,root2.Left) && checkIsSame(root1.Right,root2.Right)
}

func main() {
	arr1 := []int{5,15,3,7,13,17}
	arr2 := []int{5,15,4,11,13,17}
	arr3 := []int{5,15,4,8,13,17}
	b1 := CreateTree_5(arr1)
	b2 := CreateTree_5(arr2)
	b3 := CreateTree_5(arr3)
	fmt.Println(checkIsSame(b1.Root,b2.Root))
	fmt.Println(checkIsSame(b1.Root,b3.Root))
}