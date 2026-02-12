package main

import "fmt"

const Max int = 3

func main() {
	var a int = 10
	var ip *int
	ip = &a
	fmt.Printf("a 变量地址是：%x\n", &a)
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	var prt *int
	// 判断 prt 是否为空
	if prt != nil {
		fmt.Println("prt 值不为空")
	}
	if prt == nil {
		fmt.Println("prt 值为空")
	}

	array := []int{1, 2, 3}
	var i int
	// 指针数组
	var prtArray [Max]*int
	for i = 0; i < Max; i++ {
		//将地址放到指针数组里
		prtArray[i] = &array[i]
	}
	for i = 0; i < Max; i++ {
		//输出指针数组中指针指向的内容
		fmt.Printf("*prt[%d] 的值为: %d\n", i, *prtArray[i])
	}
}
