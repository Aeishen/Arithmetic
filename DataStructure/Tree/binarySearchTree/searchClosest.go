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

package Tree

import (
	"math"
)

func (t *MyBintree)SearchClosestValue(root *Node,target float64) int{
	rootValue := root.Data.(int)
	child := root.Left

	if t.CompareBig(rootValue,target) == -1{
		child = root.Right
	}

	if child == nil{
		return rootValue
	}else{
		childValue := t.SearchClosestValue(child,target)
		if math.Abs(float64(childValue) - target) > math.Abs(float64(rootValue) - target) {
			return rootValue
		}else{
			return childValue
		}
	}
}

