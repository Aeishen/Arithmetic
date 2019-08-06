/*
@author : Aeishen
@data :  19/08/06, 15:24
@description :sunday 算法 from https://www.jianshu.com/p/2e6eb7386cd3("简书",原文作者：houskii)
*/

/*
解释：
	Sunday算法是从前往后匹配，在匹配失败时关注的是主串中参加匹配的最末位字符的下一位字符。
	1.如果该字符没有在模式串中出现则直接跳过，即移动位数 = 模式串长度 + 1；
	2.否则，其移动位数 = 模式串长度 - 该字符最右出现的位置(以0开始) = 模式串中该字符最右出现的位置到尾部的距离 + 1。

Sunday算法的缺点
	看上去简单高效非常美好的Sunday算法，也有一些缺点。因为Sunday算法的核心依赖于move数组，而move数组的值则取决于模式串，那么就可能存在模式串构造出很差的move数组。例如下面一个例子
	主串：baaaabaaaabaaaabaaaa
	模式串：aaaaa
	这个模式串使得move[a]的值为1，即每次匹配失败时，只让模式串向后移动一位再进行匹配。这样就让Sunday算法的时间复杂度飙升到了O(m*n)，也就是字符串匹配的最坏情况

总结
	当然，也不能因为存在最坏的情况就直接否定Sunday算法，大多数情况下，Sunday依然是一个简单高效的算法，值得我们熟练学习掌握。

*/

package main

import (
	"fmt"
)

const ASCII_SIZE  =  126

func sunday(total string, part string) int {
	startIndex := 0
	len_t := len(total)
	len_p := len(part)

	moveStepsMap := make(map[int32]int)
	for _,v := range total{
		moveStepsMap[v] = len_p + 1    //主串参与匹配最末位字符移动到该位需要移动的位数
	}
	for i,v:= range part{
		moveStepsMap[v] = len_p - i
	}
	for startIndex + len_p <= len_t {
		count := 0
		for total[startIndex + count] == part[count]{
			count ++   //模式串已经匹配了的长度
			if count == len_p{
				return startIndex
			}
		}
		if startIndex + len_p == len_t {
			break
		}
		startIndex = startIndex + moveStepsMap[int32(total[startIndex + len_p])] //Sunday算法的核心
		/*
		   这是最关键的一句，涉及到Sunday算法的核心，也就是模式串在主串中的“跳跃”，我们把这句代码分解一下就好理解的多:
		    nextPos := startIndex + len_p;      //跳到startIndex + len_p，也就是模式串长度后的一个字符的位置
			c := total[nextPos];                //获取转跳后位置的字符
			moveLength := moveStepsMap[c];      //根据字符值在moveStepsMap中查找模式串需要跳跃的长度
			startIndex += moveLength;           //让模式串首部在字符串的位置加上跳跃的长度，完成跳跃
		*/
	}
	return -1
}

func main() {
	total := "substring ingsearch"
	part := "search"
	fmt.Println(sunday(total , part))
}
