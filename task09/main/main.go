package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

type student01 struct {
	Name string
	Age  int
}

func main() {
	list = make(map[string]Student)
	student := Student{Name: "张三"}
	//值拷贝过程
	list["student"] = student
	//编译失败  list["student"] 值引用的特点 是只读，所以	list["student"].Name = "李四" 不允许
	//list["student"].Name = "李四"
	fmt.Println(list["student"])

	//方法1
	//是先做一次值拷贝，做出一个tmpStudent副本,
	//然后修改该副本，
	//然后再次发生一次值拷贝复制回去，
	//list["student"] = tmpStudent,但是这种会在整体过程中发生2次结构体值拷贝，性能很差。
	tmpStudent := list["student"]
	tmpStudent.Name = "王五"
	list["student"] = tmpStudent
	fmt.Println(list["student"])

	//将map的类型的value由Student值，改成Student指针。
	//指针本身是常指针，不能修改，只读属性，但是指向的Student是可以随便修改的，而且这里并不需要值拷贝。只是一个指针的赋值。
	list2 := make(map[string]*Student)
	student2 := Student{Name: "AceId"}
	list2["student2"] = &student2
	list2["student2"].Name = "LDB"
	fmt.Println(list2["student2"].Name)

	m := make(map[string]*student01)

	//定义student数组
	stus := []student01{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//将数组依次添加到map中
	//for _, stu := range stus 循环中，stu 这个变量只在循环开始前创建了一次。
	//在每一次迭代中，Go 只是把数组当前元素的值 复制 给这个唯一的 stu 变量。
	//这意味着：
	//&stu 的地址始终不变。无论循环到第几个元素，&stu 指向的都是内存中同一个位置（即循环变量 stu 所在的栈地址）。
	for _, stu := range stus {
		// 显式创建副本 (Go 1.22+ 中这行其实有点冗余，但无害)
		newStu := stu // 关键步骤
		m[stu.Name] = &newStu
	}
	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}

}
