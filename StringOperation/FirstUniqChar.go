/*
@author : Aeishen
@data :  19/09/25, 15:48
@description :
*/

/*
题目：
    给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:
	s = "leetcode"
	返回 0.

	s = "loveleetcode",
	返回 2.

注意:
	您可以假定该字符串只包含小写字母。
*/

package main

import "fmt"

func main() {
	//s := "leetcode"
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	sMap := map[int32]int{}
	for _,v := range s{
		sMap[v] ++
	}
	for i,v := range s{
		if sMap[v] == 1{
			return i
		}
	}
	return -1
}