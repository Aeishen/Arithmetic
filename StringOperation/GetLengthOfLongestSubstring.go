/*
@author : Aeishen
@data :  19/07/31, 17:50
@description :
*/

/*
题目：
     给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
	输入: "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
	输入: "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

package main

import "fmt"

func main() {
	s := "aaasdfaasd"
	fmt.Println(lengthOfLongestSubstring(s))
}
func lengthOfLongestSubstring(s string) (int ,[]byte){
	len_s := len(s)

	if len_s <= 0{
		return 0,[]byte{}
	}

	len_diff := 0
	s_diff := []byte{}
	longest_s_diff := s_diff

	for i := 0; i< len_s; i++{
		ok,index := check(s_diff ,s[i])
		if !ok{
			s_diff = append(s_diff,s[i])
		}else{
			s_diff = append(s_diff[index + 1:],s[i])
		}

		if len_diff < len(s_diff){
			len_diff = len(s_diff)
			longest_s_diff = s_diff
		}
	}

	return  len_diff,longest_s_diff
}

func check(s []byte,s_ byte)(bool,int){
	for i,v:= range s{
		if v == s_{
			return true,i
		}
	}
	return false,0
}