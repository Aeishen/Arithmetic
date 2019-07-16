/*
@author : Aeishen
@data :  19/07/16, 14:21

@description : 创建二叉查找树
*/

package Tree

import "fmt"

type Object interface{}

type Node struct {
	Data Object
	Left *Node
	Right *Node
}

type MyBintree struct {
	Root *Node
	Size int
	CompareBig func(data1 Object,data2 Object)bool
}

func NewMyBintree() *MyBintree{
	var b MyBintree
	b.Root = &Node{}
	b.Size = 0
	b.CompareBig = nil
	return &b
}

// 查找元素
func (b *MyBintree)Search(data Object) *Node{
	if data == nil || b.Size == 0{
		return nil
	}
	if b.Root.Data == data {
		return b.Root
	}else{
		return b.search(b.Root,data)
	}
}

// 查找元素递归调用
func (b *MyBintree)search(root *Node, data Object) *Node{
	if root == nil{
		return nil
	}else if root.Data == data{
		return root
	}else if b.CompareBig(root.Data,data) {
		return b.search(root.Left, data)
	}else{
		return b.search(root.Right, data)
	}
}

// 添加元素
func (b *MyBintree)Add(data Object) bool{
	if data == nil{
		return false
	}
	newNode := &Node{data,nil,nil}
	if b.Size == 0{
		b.Root = newNode
	}else{
		b.Root = b.add(b.Root, newNode)
	}

	b.Size ++
	return true
}

// 添加元素递归调用
func (b *MyBintree)add(root *Node, newNode *Node) *Node{
	if root == nil{
		root = newNode
		return root
	}
	if b.CompareBig(root.Data,newNode.Data) {
		root.Left = b.add(root.Left, newNode)
	}else{
		root.Right = b.add(root.Right, newNode)
	}
	return root
}

// 删除元素
func (b *MyBintree)Delete(data Object) bool{
	delNode := b.Search(data)                              //获取要删除的节点p
	if delNode == nil{
		return false
	}
	if delNode.Left == nil && delNode.Right == nil{        //若删除的是叶子节点p
		delNode = nil
	}else if delNode.Left == nil || delNode.Right == nil{  //若删除的节点p只有左子树pL或右子树pR
		if delNode.Left == nil{
			fmt.Println("000",&delNode)
			delNode = delNode.Right
			fmt.Println("111",&delNode)
		}else{
			delNode = delNode.Left
		}
	}else{                                                     //若删除的节点p存在左子树pL和右子树pR
		_,leftMaxNode:= b.findLimitValueNode(delNode.Left)     //取p的左子树中最大值的节点（取p的右子树中最小值的节点也可）
		delNode.Data = leftMaxNode.Data
		leftMaxNode = nil
	}
	return true
}

// 获取该树最大值节点
func (b *MyBintree)GetMaxNode() *Node{
	if b.Size == 0{
		return nil
	}
	_, maxNode := b.findLimitValueNode(b.Root)
	return maxNode
}

// 获取该数最小值节点
func (b *MyBintree)GetMinNode() *Node{
	if b.Size == 0{
		return nil
	}
	minNode, _ := b.findLimitValueNode(b.Root)
	return minNode
}

//找某个节点下所有子节点的最小/大值的节点
func (b *MyBintree)findLimitValueNode(root *Node)(*Node,*Node){
	if root == nil {                                 // 该节点为nil
		return nil,nil
	}else if root.Left == nil && root.Right == nil{  // 该节点不存在左右子树
		return root,root
	}else{
		minNode,maxNode  := root,root
		if root.Left == nil {                        // 该节点存在右子树
			for root.Right != nil{
				maxNode = root.Right
				root = root.Right
			}
		}else if root.Right == nil {                 // 该节点存在左子树
			for root.Left != nil{
				minNode = root.Left
				root = root.Left
			}
		}else{                                       // 该节点同时存在左子树右子树
			root_ := &Node{root.Data,root.Left,root.Right}  // 复制一份该节点，防止第一个for循环影响第二个for循环
			for root.Left != nil{
				minNode = root.Left
				root = root.Left
			}
			for root_.Right != nil{
				maxNode = root_.Right
				root_ = root_.Right
			}
		}
		fmt.Println("12",&minNode)
		return minNode,maxNode
	}
}
