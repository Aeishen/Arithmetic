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

	//// 测试createBintree -----------------------------------------
	b := Tree.NewMyBintree(11)
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

	b.Add(10)
	fmt.Println(b.Size,b.Root)

	b.Add(15)
	fmt.Println(b.Size,b.Root)

	b.Add(5)
	fmt.Println(b.Size,b.Root)

	b.Add(19)
	fmt.Println(b.Size,b.Root)

	b.Add(12)
	fmt.Println(b.Size,b.Root)

	b.Add(4)
	fmt.Println(b.Size,b.Root)

	b.Add(9)
	fmt.Println(b.Size,b.Root)
	fmt.Println("获取树的所有元素的值（前序）：",b.PreorderTraversal())
	fmt.Println("获取树的所有元素的值（中序）：",b.InorderTraversal())
	fmt.Println("获取树的所有元素的值（后序）：",b.PostorderTraversal())
	fmt.Println("获取树的最大值结点",b.GetTreeMaxNode())
	fmt.Println("获取树的最小值结点",b.GetTreeMinNode())

	fmt.Println("查询111是否存在于树中：",b.Search(111))
	fmt.Println("查询15是否存在于树中：",b.Search(15))


	if b.Delete(4){
		fmt.Println("删除4后获取树的最小值结点",b.GetTreeMinNode())
		fmt.Println("删除4后获取树的size",b.Size)
		fmt.Println("删除4后获取树的所有元素的值（中序）：",b.InorderTraversal())
	}else {
		fmt.Println("树中不存在4")
	}

	if b.Delete(15){
		fmt.Println("删除15后获取树的最小值结点",b.GetTreeMinNode())
		fmt.Println("删除15后获取树的size",b.Size)
		fmt.Println("删除15后获取树的所有元素的值（中序）：",b.InorderTraversal())
	}else {
		fmt.Println("树中不存在15")
	}

	if b.Delete(19){
		fmt.Println("删除19后获取树的最大值结点",b.GetTreeMaxNode())
		fmt.Println("删除19后获取树的size",b.Size)
		fmt.Println("删除19后获取树的所有元素的值（中序）：",b.InorderTraversal())
	}else {
		fmt.Println("树中不存在19")
	}


	if b.Delete(100){
		fmt.Println("删除100后获取树的最小值结点",b.GetTreeMinNode())
		fmt.Println("删除100后获取树的size",b.Size)
		fmt.Println("删除100后获取树的所有元素的值（中序）：",b.InorderTraversal())
	}else {
		fmt.Println("树中不存在100")
	}

	//// 测试searchClosest --------------------
    fmt.Println(b.SearchClosestValue(b.Root,1100.6))
}