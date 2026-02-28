package main

import (
	"fmt"
	"reflect"
)

func main() {

	//结构体比较
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	//sn3:= struct {
	//	name string
	//	age  int
	//}{age:11, name:"qq"}

	// sn1 与 sn3 属性顺序不同，所以就是不同的结构体了，所以不能比较
	//if sn1 == sn3 {
	//	fmt.Println("sn1 == sn2")
	//}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice，则结构体不能用==比较。
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
	//可以使用reflect.DeepEqual进行比较
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	} else {
		fmt.Println("sm1 != sm2")
	}

	//切片追加, make初始化均为0
	s := make([]int, 10)
	s = append(s, 1, 2, 3)
	//[0 0 0 0 0 0 0 0 0 0 1 2 3]
	fmt.Println(s)
	//slice 拼接问题
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//两个slice在append的时候，
	//记住需要进行将第二个slice进行...打散再拼接。
	//s1 = append(s1,s2)
	s1 = append(s1, s2...)
	fmt.Println(s1)

	//new和make的区别：
	//二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；
	//而new用于类型的内存分配，并且内存置为零。所以在我们编写程序的时候，就可以根据自己的需要很好的选择了。
	//make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
	//list := new([]int)
	//list = append(list, 1)
	//fmt.Println(list)

}

// 函数返回值问题,如果返回值 指定名称，那么所有返回值都要指定名称
//
//	func myFunc(x,y int) (sum int , error){
//		return x+y,nil
//	}
//
// 正确写法
func myFunc(x, y int) (sum int, err error) {
	return x + y, nil
}

// GetValue string与nil类型
func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	//nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
	//return nil, false
	return "不存在数据", false
}
