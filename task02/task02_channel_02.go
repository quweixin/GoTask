package task02

/**
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/
import (
	"fmt"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产：%d\n", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("消费：%d\n", v)
	}
	fmt.Println("所有数据已消费完毕。")
}

func TestProducerConsumer() {
	ch := make(chan int, 6)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}
