/*
@author : Aeishen
@data :  19/08/07, 10:28

@description : LeetCode 上第 10 号问题
*/

/*
题目：
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和'*' 的正则表达式匹配。（所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。）
    '.' 匹配任意单个字符
    '*' 匹配零个或多个前面的那一个元素

说明：
    s 可能为空，且只包含从a-z 的小写字母。
	p 可能为空，且只包含从a-z 的小写字母，以及字符 .和 *。

示例 1:
	输入: s = "aa"；p = "a"
	输出: false
	解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
	输入:s = "aa"；p = "a*"
	输出: true
	解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了

示例 3:
	输入:s = "ab"；p = ".*"
	输出: true
	解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4:
	输入:s = "aab"；p = "c*a*b"
	输出: true
	解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

示例 5:
	输入:s = "mississippi"；p = "mis*is*p*."
	输出: false
*/

package main

import "fmt"

// 递归暴力解法：字符串可以划分成为一个个字符，这样字符串比较问题就会变成字符比较问题，这样一来，我们就可以把
//      问题看成，决定 s[i,…n] 是否能够匹配 p[j,…m] 的条件是在s[i] 和 p[j] 匹配情况下，子问题 s[i+1,…n]
//      能不能够匹配 p[j+1,…m]

func isMatch(s string,p string)bool{
	if s == p {
		return true
	}

	isFirstMatch := false
	if (len(s) > 0 && len(p) > 0) && (s[0] == p[0] || string(p[0]) == "."){
		isFirstMatch = true
	}

	if len(p) > 1 && string(p[1]) == "*"{
		return (isFirstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}

	return isFirstMatch && isMatch(s[1:],p[1:])
}

/*
	上面的实现之所以被称为暴力求解是因为子问题的答案没有被记录，也就是说如果当前要用到之前的子问题的答案，
     我们还得去计算之前计算过的子问题。
*/


func main() {
	fmt.Println(isMatch( "aa","a"))
	fmt.Println(isMatch( "aa","a*"))
	fmt.Println(isMatch( "ab",".*"))
	fmt.Println(isMatch( "aab","c*a*b"))
	fmt.Println(isMatch( "mississippi","mis*is*p*."))
}
