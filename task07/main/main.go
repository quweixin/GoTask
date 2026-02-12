package main

import "fmt"

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
}
