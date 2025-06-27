package task02

/**
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
import (
	"fmt"
	"sync"
)

func increment(num *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	for i := 0; i < 1000; i++ {
		*num += 1
	}
	mu.Unlock()
}

func TestIncrement() {
	num := 0
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go increment(&num, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(num)
}
