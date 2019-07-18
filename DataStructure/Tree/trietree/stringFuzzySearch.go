/*
@author : Aeishen
@data :  19/07/18, 10:40

@description :
*/

/*
题目：
	Design a data structure that supports the following two operations:
		void addWord(word)
		bool search(word)
		search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

	For example:
		addWord("bad")
		addWord("dad")
		addWord("mad")
		search("pad") -> false
		search("bad") -> true
		search(".ad") -> true
		search("b..") -> true

	Note:
		You may assume that all words are consist of lowercase letters a-z.
*/

/*
解题思路：
	使用Trietree实现
*/
package Tree


type Node_1 struct {
	isEnd bool
	Children map[byte]*Node_1
}

type MyStructure struct {
	Root *Node_1
}

func newMyNode_1() *Node_1{
	return &Node_1{false,make(map[byte]*Node_1)}
}

func NewMyStructure() *MyStructure{
	return &MyStructure{newMyNode_1()}
}

func (this *MyStructure)AddWord(word string)  {
	curRoot := this.Root

	for i:= 0;i<len(word);i++{
		if _,ok := curRoot.Children[word[i]];!ok{
			curRoot.Children[word[i]] = newMyNode_1()
		}
		curRoot = curRoot.Children[word[i]]
	}
	curRoot.isEnd = true
}


func (this *MyStructure)Search(word string) bool {
	return this.search(word, this.Root, 0)
}

func (this *MyStructure)search(word string, curRoot *Node_1, index int) bool {
	if index == len(word) - 1 && word[index] != '.'{
		if _,ok := curRoot.Children[word[index]];ok{
			return curRoot.Children[word[index]].isEnd
		}else{
			return false
		}
	}else if index == len(word) - 1 && word[index] == '.'{
		return true
	}

	if len(curRoot.Children) > 0{
		if word[index] != '.' {
			if child,ok := curRoot.Children[word[index]]; ok{
				return this.search(word, child, index+1)
			}
			return false
		}else{
			for _,child := range curRoot.Children{
				if this.search(word, child, index+1){
					return true
				}
			}
		}
	}
	return false
}

