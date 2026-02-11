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

	//使用sort.Slice 回调
	numbers := []int{101, 5, 13, 4, 25}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("排序后的结果：", numbers)

	//c:=counter()
	//// 调用 counter，得到一个闭包函数
	//for i := 0; i < 3; i++ {
	//	//这里会输出 1，2，3
	//	fmt.Println(c())
	//	//这里会输出 1，1，1
	//	//fmt.Println(counter()())
	//}

	//| 特性 | 说明 |
	//|------|------|
	//| 捕获外部变量 | 内部函数可以访问外层函数的局部变量 |
	//| 变量持久化 | 即使外层函数返回，被捕获的变量依然存在 |
	//| 状态保持 | 多次调用闭包，可以累积或修改状态（如计数器） |
	//| 每个闭包独立 | 每次调用 `counter()` 都会创建新的 `count` 和新闭包 |

	//| 每个闭包独立 | 每次调用 `counter()` 都会创建新的 `count` 和新闭包 |
	//闭包让函数“带着环境一起走”——即使创建它的函数已经结束，它依然能记住并操作那些变量。
	c1 := counter()
	for i := 0; i < 3; i++ {
		fmt.Println(c1())
	}
	c2 := counter()
	for i := 0; i < 3; i++ {
		fmt.Println(c2())
	}

}

func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 闭包
func counter() func() int {
	count := 0
	return func() int {
		//引用了外部的 count 变量
		//虽然 counter() 执行完后“按理说” count 应该被销毁，
		//但因为闭包还在使用它，Go 的运行时会自动将 count 分配在堆上（而不是栈上），让它继续存活！
		count++
		return count
	}
}
