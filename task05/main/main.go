package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func main() {
	result := func(x, y int) int {
		return x + y
	}(3, 4)
	fmt.Println("结果：", result)

	//匿名函数
	var f = func() {
		fmt.Println("hello world")
	}
	f()

	//回调函数
	multiply := func(a, b int) int {
		return a * b
	}
	callBackFunc := operate(3, 4, multiply)
	fmt.Println("结果：", callBackFunc)

	//匿名函数 回调函数
	callBackFunc = operate(3, 4, func(a, b int) int {
		return a + b
	})
	fmt.Println("结果：", callBackFunc)

	//使用channel 方式，防止go 协程退出
	done := make(chan bool)
	go func(msg string) {
		fmt.Println("Goroutine:", msg)
		done <- true // 发送信号表示完成
	}("Hello from anonymous func!")
	<-done

	//使用 sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(msg string) {
			//defer 是无论如何都要执行，相当于Java的finally
			defer wg.Done()
			fmt.Println(msg)
		}("Goroutine: Hello from anonymous func! " + strconv.Itoa(i))
	}
	wg.Wait() // 阻塞，直到所有 Done 被调用

	numbers := []int{101, 5, 13, 4, 25}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("排序后的结果：", numbers)
}

func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}
