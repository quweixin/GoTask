package main

import (
	"GoTask/task01"
	"fmt"
)

func main() {
	// task01_出现一次的数字
	var numbs = []int{4, 1, 2, 1, 2}
	result := task01.SingleNumber(numbs)
	fmt.Println(result)
	// task01_回文
	//var s = "A man, a plan, a canal: Panama"
	//var s = "madam"
	var s = "上海自来水来自海上"
	var isPalindrome = task01.IsPalindrome(s)
	if isPalindrome {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}

	// task01_有效括号
	//var str = "()[]{}"
	var str = "([)]"
	isValid := task01.IsValid(str)
	if isValid {
		fmt.Println("有效括号")
	} else {
		fmt.Println("无效括号")
	}

	// 最长公共前缀
	var stirs = []string{"flower", "flow", "flight"}
	prefix := task01.LongestCommonPrefix(stirs)
	fmt.Println("最长公共前缀：" + prefix)
}
