/*
@author : Aeishen
@data :  19/07/16, 14:21

@description : 创建二叉查找树
*/

package Tree

type Object interface{}

type Node struct {
	Data Object
	Left *Node
	Right *Node
}

type MyBintree struct {
	Root *Node
	Size int
	CompareBig func(data1 Object,data2 Object)int
}

func newNode(data Object) *Node{
	return &Node{data,nil,nil}
}

func NewMyBintree(root Object) *MyBintree{
	return &MyBintree{newNode(root),1,nil}
}

// 查找元素
func (b *MyBintree)Search(data Object) bool{
	if data == nil || b.Size == 0{
		return false
	}
	return b.search(b.Root,data) != nil
}

// 在以root为根节点的树中查找值为data元素是否存在
func (b *MyBintree)search(root *Node, data Object) *Node{
	if root == nil{
		return nil
	}else if b.CompareBig(root.Data,data) == 0{
		return root
	}else if b.CompareBig(root.Data,data) == 1 {
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
	newNode := newNode(data)
	if b.Size == 0{
		b.Root = newNode
	}else{
		b.Root = b.add(b.Root, newNode)
	}

	b.Size ++
	return true
}

// 将值为data的节点添加以root为根节点的树中
func (b *MyBintree)add(root *Node, newNode *Node) *Node{
	if root == nil{
		root = newNode
		return root
	}
	if b.CompareBig(root.Data,newNode.Data) == 1 {
		root.Left = b.add(root.Left, newNode)
	}else{
		root.Right = b.add(root.Right, newNode)
	}
	return root
}

// 删除元素
func (b *MyBintree)Delete(data Object) bool{
	if data == nil || !b.Search(data)  {
		return false
	}
	b.Root = b.delete(b.Root,data)
	b.Size --
	return true
}

// 删除以root为根节点的树中值为data的节点
func (b *MyBintree)delete(root *Node, data Object)(*Node) {
	if root == nil{
		return root
	}
	compareResult := b.CompareBig(root.Data, data)
	if compareResult == 1{
		root.Left = b.delete(root.Left, data)
	}else if compareResult == -1{
		root.Right = b.delete(root.Right, data)
	}else if compareResult == 0 {
		if root.Right == nil && root.Left == nil {
			root = nil
		} else if root.Right == nil {
			root = root.Left
		} else if root.Left == nil {
			root = root.Right
		} else {
			root.Data = b.getMaxNodeByNode(root.Left).Data
			root.Left = b.delete(root.Left, root.Data)
		}
	}
	return root
}

// 获取该树最大值
func (b *MyBintree)GetTreeMaxNode() *Node{
	return b.getMaxNodeByNode(b.Root)
}

// 获取该树最小值
func (b *MyBintree)GetTreeMinNode() *Node{
	return b.getMinNodeByNode(b.Root)
}

//找到以root为根节点的树中最小值的节点
func (b *MyBintree)getMinNodeByNode(root *Node) *Node{
	for root.Left != nil{
		root = root.Left
	}
	return root
}

//找到以root为根节点的树中最大值的节点
func (b *MyBintree)getMaxNodeByNode(root *Node) *Node{
	for root.Right != nil{
		root = root.Right
	}
	return root
}

//前序遍历递归实现
func (b *MyBintree)PreorderTraversal() (all []Object){
	return b.getAllData(all,b.Root,-1)
}

//中序遍历递归实现
func (b *MyBintree)InorderTraversal() (all []Object){
	return b.getAllData(all,b.Root,0)
}

//后序遍历递归实现
func (b *MyBintree)PostorderTraversal() (all []Object){
	return b.getAllData(all,b.Root,1)
}

//获取以root为根节点的树中的所有节点的值
func (b *MyBintree)getAllData(all []Object,root *Node,tag int) []Object{
	//前序遍历：先访问当前节点，再依次递归访问左右子树
	if root != nil && tag == -1{
		all = append(all,root.Data)
		all = b.getAllData(all,root.Left,-1)
		all = b.getAllData(all,root.Right,-1)
	}
	//中序遍历：先递归访问左子树，再访问自身，再递归访问右子树
	if root != nil && tag == 0{
		all = b.getAllData(all,root.Left,0)
		all = append(all,root.Data)
		all = b.getAllData(all,root.Right,0)
	}
	//后序遍历：先递归访问左右子树，最后再访问当前节点。
	if root != nil && tag == 1{
		all = b.getAllData(all,root.Left,1)
		all = b.getAllData(all,root.Right,1)
		all = append(all,root.Data)
	}

	return all
}

//层次遍历（广度优先遍历）
func (b *MyBintree)LevelOrderTraversal() (all []Object){
	nodes := []*Node{b.Root}

	for len(nodes) > 0{
		curNode := nodes[0]
		nodes = nodes[1:]
		all = append(all,curNode.Data)
		if curNode.Left != nil{
			nodes = append(nodes,curNode.Left)
		}
		if curNode.Right != nil{
			nodes = append(nodes,curNode.Right)
		}
	}
	return
}

// 求树的深度
func (b *MyBintree)GetTreeDepth(root *Node) int{
	if root == nil {
		return 0
	}
	countL := b.GetTreeDepth(root.Left) + 1
	countR := b.GetTreeDepth(root.Right) + 1
	if countL > countR{
		return countL
	}
	return countR
}

//获取树第k层的节点个数
func (b *MyBintree)GetNodeCount(root *Node,k int) int{
	if root == nil || k < 1 || k > b.GetTreeDepth(b.Root) {
		return 0
	}
	if k == 1{
		return 1
	}
	return b.GetNodeCount(root.Left,k-1) + b.GetNodeCount(root.Right,k-1)
}