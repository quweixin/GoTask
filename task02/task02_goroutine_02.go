package task02

import (
	"fmt"
	"sync"
	"time"
)

/**
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

func TaskA() {
	time.Sleep(1 * time.Second)
	fmt.Println("taskA 执行完毕")
}
func TaskB() {
	time.Sleep(2 * time.Second)
	fmt.Println("taskB 执行完毕")
}
func TaskC() {
	time.Sleep(3 * time.Second)
	fmt.Println("taskC 执行完毕")
}

func TaskScheduler(tasks map[string]func()) map[string]time.Duration {
	results := make(map[string]time.Duration)
	//sync.WaitGroup 用于等待多个 Goroutine 完成。
	var wg sync.WaitGroup
	// 锁机制，用于保护共享变量
	mu := sync.Mutex{}
	for name, task := range tasks {
		// 计数器 加1
		wg.Add(1)
		go func(name string, task func()) {
			// 在 goroutine 完成时减少计数器
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			// 使用锁机制保护共享变量
			mu.Lock()
			results[name] = duration
			// 解锁
			mu.Unlock()
		}(name, task)
	}
	// // 等待所有 goroutine 完成
	wg.Wait()
	return results
}
