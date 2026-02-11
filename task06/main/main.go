package main

import (
	"fmt"
)

func main() {
	//x:=cmp.Compare(1,2)
	//fmt.Println(x)
	PrintAny(42)      // Value: 42, Type: int
	PrintAny("hello") // Value: hello, Type: string
	PrintAny(3.14)    // Value: 3.14, Type: float64

	numbers := []int{1, 2, 3, 4, 5}
	index := FindIndex(numbers, 3)
	fmt.Println("Index of 3:", index)

	names := []string{"Alice", "Bob", "Charlie", "David"}
	index = FindIndex(names, "Charlie")
	fmt.Println("Index of Charlie:", index)

	// Number 联合约束
	fmt.Println("Add(1, 2):", Add(1, 2))
	fmt.Println("Add(3.14, 2.718):", Add(3.14, 2.718))

	// Stringer 约束
	person := Person{Name: "Alice", Age: 30}
	PrintString(person)
}

// PrintAny any 约束，any 是空接口 interface{} 的别名，表示任何类型都可以。
func PrintAny[T any](value T) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// FindIndex comparable 约束 comparable 表示类型支持 == 和 != 操作符。
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// Number 联合约束
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

type Stringer interface {
	String() string
}

// PrintString 自定义约束
func PrintString[T Stringer](v T) {
	fmt.Println(v.String())
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// NumericStringer 要求类型是数字且实现 String() 方法
type NumericStringer interface {
	Number
	String() string
}
