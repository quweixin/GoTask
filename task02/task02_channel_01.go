package task02

import (
	"fmt"
	"sync"
)

func writeNumber(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func readNumber(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range c {
		fmt.Println(num)
	}
}

func RunCommunication() {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(2)
	go writeNumber(c, &wg)
	go readNumber(c, &wg)
	wg.Wait()
}
