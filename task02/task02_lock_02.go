package task02

import (
	"sync"
	"sync/atomic"
)

/**
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

func atomicIncrement(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func TestAtomicIncrement() {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go atomicIncrement(&counter, &wg)
	}
	wg.Wait()
	println(counter)
}
