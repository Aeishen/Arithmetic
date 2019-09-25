/*
@author : Aeishen
@data :  19/09/25, 17:25
@description :
*/

/*
题目：
    编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
示例 1:
	输入: ["flower","flow","flight"]
	输出: "fl"
示例 2:
	输入: ["dog","racecar","car"]
	输出: ""
解释:
	输入不存在公共前缀。
说明:
	所有输入只包含小写字母 a-z 。
*/
package main

import "fmt"

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i:= 1;i<len(strs); i++{
		res = getCommonPrefix(res,strs[i])
		if res == ""{
			return ""
		}
	}
	return res
}

func getCommonPrefix(s1 string,s2 string)string{
	l := len(s1)
	if l > len(s2){
		l = len(s2)
	}

	for i := 0;i < l;i++{
		if s1[i] != s2[i] {
			if i>0{
				return s1[:i]
			}
			return ""
		}
	}
	return s1[:l]
}
