/*
@author : Aeishen
@data :  19/09/26, 18:07
@description :
*/

/*
题目：
     给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
	输入: "A man, a plan, a canal: Panama"
	输出: true
示例 2:
	输入: "race a car"
	输出: false
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome("race a car"))
	//fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("`l;`` 1o1 ??;l`"))
}

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	j := len(s)-1
	for i:=0;i<len(s);i++{
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			for k := j;k>=0;k--{
				if (s[k] >= 'A' && s[k] <= 'Z') || (s[k] >= '0' && s[k] <= '9') {
					if s[k] == s[i]{
						j = k - 1
						break
					}else{
						return false
					}
				}
			}
		}
	}
	return true
}
