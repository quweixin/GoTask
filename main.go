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
}
