/*
   @File : WordBreak
   @Author: Aeishen
   @Date: 2020/1/8 23:38
   @Description:139. 单词拆分
*/

/*
题目：
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

示例 1：
	输入: s = "leetcode", wordDict = ["leet", "code"]
	输出: true
	解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。


示例 2：
	输入: s = "applepenapple", wordDict = ["apple", "pen"]
	输出: true
	解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。注意你可以重复使用字典中的单词。

示例 3：
	输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
	输出: false

说明:
	拆分时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。。

思路：
   暴力：最简单的方法就是用递归和回溯，为了找到解，我们检查字典单词中每个单词的可能前缀，如果在字典中出现过则去掉这个前缀后剩余的
        部分回归调用，同时，如果某次函数调用中发现整个字符串都已经被拆分且在字典中出现过，则返回true
        时间复杂度：O(N)

*/

package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak1(s , wordDict))
}

func wordBreak1(s string, wordDict []string) bool {
	wordMap := make(map[string]int)
	for _,v := range wordDict{
		wordMap[v] =1
	}
	return violence(s, wordMap, 0)
}

func violence(s string, wordMap map[string]int, start int) bool  {
    if start == len(s){
		return true
	}

	for i := start + 1;i <= len(s); i++{
		curWord := s[start:i]
		if wordMap[curWord] == 1 && violence(s, wordMap, i){
			return true
		}
	}

	return false

}