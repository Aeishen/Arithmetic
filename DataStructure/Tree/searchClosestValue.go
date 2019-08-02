/*
@author : Aeishen
@data :  19/07/24, 17:37

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
	"math"
)

//创建二叉树
func CreateTree_6() *Tree.MyBintree{
	b := Tree.NewMyBintree(10)
	b.CompareBig = func(data1 Tree.Object, data2 Tree.Object) int {
		switch data2.(type) {
		case float64:
			v := float64(data1.(int)) - data2.(float64)
			if v > 0{
				return 1
			}else if v < 0 {
				return -1
			}
		case int:
			if data1.(int) >  data2.(int){
				return 1
			}else if data1.(int) <  data2.(int){
				return -1
			}else{
				return 0
			}
		}
		return 0
	}
	arr := []int{5,15,3,7,13,17}

	for _,v := range arr{
		b.Add(v)
	}
	return b
}

func search(b *Tree.MyBintree,root *Tree.Node, target float64)Tree.Object{
	rootValue := root.Data
	child := root.Left
	if b.CompareBig(rootValue,target) <0{
		child = root.Right
	}

	if child == nil{
		return rootValue
	}else{
		childValue := search(b ,child, target)
		if math.Abs(float64(childValue.(int)) - target)>math.Abs(float64(rootValue.(int)) - target ){
			return rootValue
		}else{
			return childValue
		}
	}
}

func main() {
	b:= CreateTree_6()
	fmt.Println(search(b,b.Root,3.99999))
}