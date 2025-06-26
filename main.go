package main

import (
	"GoTask/task01"
	"fmt"
)

func main() {
	var numbs = []int{4, 1, 2, 1, 2}
	result := task01.SingleNumber(numbs)
	fmt.Println(result)
}
