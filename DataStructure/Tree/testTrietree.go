/*
@author : Aeishen
@data :  19/07/17, 17:22

@description : 测试Trie树
*/

package main

import (
	"./trietree"
	"fmt"
)

func main() {

	//// test createTrieTree---------------------------
	t := Tree.NewMyTrieTree()
	t.Add("code")
	//t.Add("cook")
	//t.Add("five")
	//t.Add("file")
	//t.Add("fat")
	//t.Add("hi")
	//t.Add("player")
	//t.Add("play")

	//// 使用Search查询存在的整个字符串
	//fmt.Println(t.Search("code"))
	//fmt.Println(t.Search("five"))
	//fmt.Println(t.Search("hi"))
	//// 使用Search查询存在的字符串的前缀
	//fmt.Println(t.Search("cod"))
	//fmt.Println(t.Search("fiv"))
	//fmt.Println(t.Search("h"))
	//// 使用StartWith查询存在的整个字符串
	//fmt.Println(t.StartWith("cook"))
	//fmt.Println(t.StartWith("file"))
	//fmt.Println(t.StartWith("hi"))
	//// 使用StartWith查询存在的字符串的前缀
	//fmt.Println(t.StartWith("cod"))
	//fmt.Println(t.StartWith("fi"))
	//fmt.Println(t.StartWith("h"))
	//
	//fmt.Println(t.Search("play"))
	//fmt.Println(t.StartWith("play"))
	//fmt.Println(t.Search("player"))
	//fmt.Println(t.StartWith("player"))


	//// test stringFuzzySearch------------------------
	t1 := Tree.NewMyStructure()
	t1.AddWord("bad")
	//t1.AddWord("dad")
	//t1.AddWord("mab")
	//t1.AddWord("mabdb")
    //fmt.Println(t1.Search("pad"))
	//fmt.Println(t1.Search("bad"))
	//fmt.Println(t1.Search(".ad"))
	//fmt.Println(t1.Search("b.."))
	//fmt.Println(t1.Search(".."))
	//fmt.Println(t1.Search(".a."))
	//fmt.Println(t1.Search("...d."))
	//fmt.Println(t1.Search(".a.d."))

	fmt.Println(1-0.2)
}
