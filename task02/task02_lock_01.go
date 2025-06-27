package task02

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
