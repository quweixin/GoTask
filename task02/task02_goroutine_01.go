package task02

/**
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/

func PrintOddNum(nums []int, c chan int) {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			c <- nums[i]
		}
	}
	close(c)
}

func PrintEvenNum(nums []int, c chan int) {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			c <- nums[i]
		}
	}
	close(c)
}
