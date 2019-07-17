/*
@author : Aeishen
@data :  19/07/17, 16:21

@description : 创建Trie树
*/
package Tree

import "fmt"

type Object interface {}

type Node struct {
	Children map[Object]*Node
	isEnd bool
}

type TrieTree struct {
	Root *Node
}

//构造节点
func newMyNode() *Node{
	return &Node{make(map[Object]*Node),false}
}

//构造TrieTree
func NewMyTrieTree() *TrieTree{
	return &TrieTree{newMyNode()}
}

//添加一个字符串
func (t *TrieTree)Add(s Object){

	// 获取根节点
	curRoot := t.Root

	for i := 0; i < len(s.(string)); i++ {
		if _, ok := curRoot.Children[s.(string)[i]]; !ok{
			curRoot.Children[s.(string)[i]] = newMyNode()
		}
		curRoot = curRoot.Children[s.(string)[i]]
	}

	// 最后一个节点代表一个字符串结束
	curRoot.isEnd = true
}


//查询一个字符串
func (t *TrieTree)Search(s Object)bool{

	// 获取根节点
	curRoot := t.Root

	for i := 0; i < len(s.(string)); i++ {

		if _, ok := curRoot.Children[s.(string)[i]]; !ok{
			return false
		}
		curRoot = curRoot.Children[s.(string)[i]]
	}

	return curRoot.isEnd
}

//查询一个字符串前缀是否存在
func (t *TrieTree)StartWith(s Object)bool{

	// 获取根节点
	curRoot := t.Root

	for i := 0; i < len(s.(string)); i++ {
		if _, ok := curRoot.Children[s.(string)[i]]; !ok{
			return false
		}
		curRoot = curRoot.Children[s.(string)[i]]
	}

	return true
}

//删除一个字符串
func (t *TrieTree)Delete(s Object){

	// 获取根节点
	curRoot := t.Root
    var path *Node
	for i := 0; i < len(s.(string)); i++ {
		fmt.Println(string(s.(string)[i]))
		if _, ok := curRoot.Children[s.(string)[i]]; !ok{
			return
		}
		curRoot = curRoot.Children[s.(string)[i]]
		if len(curRoot.Children) > 1{
			path = curRoot
		}
	}

	if curRoot.isEnd{
		if len(curRoot.Children) > 0{  // 删除的是字符串前缀（eg:play）
			curRoot.isEnd = false
		}else{                         // 删除的是整个字符串（eg:player）
			curRoot = nil
			fmt.Println("11",s)
			s = s.(string)[:len(s.(string)) - 1]
			fmt.Println("22",s)
			t.Delete(s)
		}
	}
}
