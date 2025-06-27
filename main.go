package main

import (
	"GoTask/task02"
	"fmt"
)

//func main() {
//	// task01_出现一次的数字
//	var numbs = []int{4, 1, 2, 1, 2}
//	result := task01.SingleNumber(numbs)
//	fmt.Println(result)
//	// task01_回文
//	//var s = "A man, a plan, a canal: Panama"
//	//var s = "madam"
//	var s = "上海自来水来自海上"
//	var isPalindrome = task01.IsPalindrome(s)
//	if isPalindrome {
//		fmt.Println("是回文")
//	} else {
//		fmt.Println("不是回文")
//	}
//
//	// task01_有效括号
//	//var str = "()[]{}"
//	var str = "([)]"
//	isValid := task01.IsValid(str)
//	if isValid {
//		fmt.Println("有效括号")
//	} else {
//		fmt.Println("无效括号")
//	}
//
//	// 最长公共前缀
//	var stirs = []string{"flower", "flow", "flight"}
//	prefix := task01.LongestCommonPrefix(stirs)
//	fmt.Println("最长公共前缀：" + prefix)
//
//	//  加一
//	//var digits = []int{1, 2, 3}
//	//var digits = []int{9}
//	var digits = []int{9, 9, 9, 9}
//	plusOneResult := task01.PlusOne(digits)
//	fmt.Println("加一后的结果：", plusOneResult)
//
//	//删除有序数组中的重复项
//	var nums = []int{1, 1, 2}
//	removeDuplicatesResult := task01.RemoveDuplicates(nums)
//	fmt.Println("删除有序数组中的重复项后的结果：", removeDuplicatesResult)
//
//	//两数之和
//	var nums03 = []int{2, 7, 11, 15}
//	target := 9
//	twoSumResult := task01.TwoSum(nums03, target)
//	fmt.Println("两数之和的结果：", twoSumResult)
//}

func main() {
	value := 20
	fmt.Println("调用函数前的值:", value)
	task02.AddTen(&value)
	// 输出修改后的值
	fmt.Println("调用函数后的值:", value)

	// 创建一个整数切片
	numbers := []int{1, 2, 3, 4, 5}

	// 输出原始切片
	fmt.Println("原始切片:", numbers)

	// 调用函数，传入切片的指针
	task02.DoubleSlice(&numbers)

	// 输出修改后的切片
	fmt.Println("修改后的切片:", numbers)

	fmt.Println("task02_goroutine_01")
	nums := []int{2, 7, 4, 6, 11, 15, 20, 21, 25}
	ch := make(chan int, 20)
	ch02 := make(chan int, 20)
	go task02.PrintOddNum(nums, ch)
	go task02.PrintEvenNum(nums, ch02)
	fmt.Println("打印奇数")
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("打印偶数")
	for i := range ch02 {
		fmt.Println(i)
	}

	tasks := map[string]func(){
		"TaskA": task02.TaskA,
		"TaskB": task02.TaskB,
		"TaskC": task02.TaskC,
	}
	fmt.Println("开始执行任务...")
	results := task02.TaskScheduler(tasks)
	fmt.Println("所有任务已完成！")
	fmt.Println("\n执行时间统计:")
	for name, duration := range results {
		fmt.Printf("%s: %.2f 秒 \n", name, duration.Seconds())
	}

	fmt.Println("面向对象01")
	var rectangle = task02.Rectangle{Width: 5.0, Height: 10.0}
	var s task02.Shape = rectangle
	fmt.Println("面积:", s.Area())
	fmt.Println("周长:", s.Perimeter())

	var circle = task02.Circle{Radius: 5.0}
	s = circle
	fmt.Println("面积:", s.Area())
	fmt.Println("周长:", s.Perimeter())

}
