/*
@author : Aeishen
@data :  19/07/16, 15:43

@description : 测试二叉查找树
*/
package main

import (
	"./bintree"
	"fmt"
)

func main() {
	b := Tree.NewMyBintree()
	b.CompareBig = func(data1 Tree.Object, data2 Tree.Object) bool {
		if data1.(int) >  data2.(int){
			return true
		}
		return false
	}

	b.Add(10)
	fmt.Println(b.Size,b.Root)

	b.Add(15)
	fmt.Println(b.Size,b.Root)

	b.Add(19)
	fmt.Println(b.Size,b.Root)

	b.Add(12)
	fmt.Println(b.Size,b.Root)

	b.Add(4)
	fmt.Println(b.Size,b.Root)

	b.Add(9)
	fmt.Println(b.Size,b.Root)


	a := b.Search(111)
	fmt.Println(a)

	a = b.Search(15)
	fmt.Println(a)

	fmt.Println(b.GetMinNode())
	fmt.Println(b.GetMaxNode())

	isDelte := b.Delete(4)
	fmt.Println(isDelte,b.GetMinNode())
}