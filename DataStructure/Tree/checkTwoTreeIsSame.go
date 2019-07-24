/*
@author : Aeishen
@data :  19/07/24, 17:06

@description :
*/

/*
题目：
	判断两棵二叉树是否结构相同,不考虑数据内容。结构相同意味着对应的左子树和对应的右子树都结构相同。
*/

/*
解题思路：
     递归
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