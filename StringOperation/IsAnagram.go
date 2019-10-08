/*
@author : Aeishen
@data :  19/10/08, 22:05
@description :
*/

/*
题目：
    给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:
	输入: s = "anagram", t = "nagaram"
	输出: true
示例 2:
	输入: s = "rat", t = "car"
	输出: false
说明:
	你可以假设字符串只包含小写字母。
进阶:
	如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

package main

import "fmt"

func main() {
	fmt.Println(isAnagram("aaaccc", "cccaa"))
}


func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
    s_map := make(map[int32]int)
	t_map := make(map[int32]int)

	for _,vs := range s{
		s_map[vs] ++
	}
	for _,vt := range t{
		t_map[vt] ++
	}

	for i,v := range s_map{
		if t_map[i] != v{
			return false
		}
	}

	return true
}