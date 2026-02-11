package main

import (
	"GoTask/task02"
)

// import (
//
//	"GoTask/task02"
//	"fmt"
//
// )
//
//	func main() {
//		var rectangle = task02.Rectangle{Width: 5.0, Height: 10.0}
//		var s task02.Shape = rectangle
//		fmt.Println("面积:", s.Area())
//		fmt.Println("周长:", s.Perimeter())
//
//		var rectangle2 = task02.Rectangle{Width: 5.0, Height: 10.0}
//		var test task02.Test = rectangle2
//		fmt.Println("Test:", test.Test())
//		fmt.Println("Test2:", test.Test2())
//
//		//var i interface{} = "hello"
//		//str  := i.(string)
//		//fmt.Println("str:", str)
//
//		//类型断言
//		var i2 interface{ } = 42
//		str,ok:= i2.(string)
//		if ok {
//			fmt.Println("str:", str)
//		}else{
//			fmt.Println("i2 is not a string")
//		}
//		//类型选择
//		printType( 42)
//		printType("hello")
//		printType(3.14)
//		printType([]int{1,2,3})
//
//		//接口组合
//		var rw ReaderWriter = File{}
//		fmt.Println("rw:", rw.Read())
//		rw.Write("hello")
//
//		//动态值 动态类型
//		//var v interface{} = 42
//		var v interface{} = "hello"
//		fmt.Printf("Type: %T, Value:%v\n", v, v)
//
//		//接口的零值
//		var i interface{}
//		fmt.Println("i:", i)
//
//		//接口demo
//		var phone Phone
//		phone = new (NokiaPhone)
//		 phone.call()
//
//		phone = new (IPhone)
//		 phone.call()
//	}
//
//	func printType ( val interface{}) {
//		switch v:=val.(type) {
//		case int:
//			fmt.Println("int:",v)
//		case string:
//			fmt.Println("string:",v)
//		case float64:
//			fmt.Println("float64:",v)
//		default:
//			fmt.Println("unknown")
//		}
//	}
func main() {
	//接口组合
	//var rw task02.ReaderWriter = task02.File{}
	//fmt.Println("rw:", rw.Read())
	//rw.Write("hello")

	//接口断言
	//var file interface{} = task02.File{Name: "张三"}
	//var rw2 task02.ReaderWriter = file
	//readerWriter,ok:= rw2.(task02.ReaderWriter)
	//if ok {
	//	rw2.Write(file.Name)
	//	fmt.Println("rw2 is readerWriter:",readerWriter)
	//	fmt.Println("rw2 is readerWriter:",readerWriter.Read())
	//}else{
	//	fmt.Println("rw2 is not readerWriter")
	//}

	//f, ok := file.(task02.File)
	//if ok {
	//	fmt.Println("f is file:",f.Name)
	//}else{
	//	fmt.Println("f is not file")
	//}
	//
	//var test interface{} = 42
	//value,ok:= test.(int)
	//if ok {
	//	fmt.Println("test is int:",value)
	//}else{
	//	fmt.Println("test is not int")
	//}

	var reader task02.Reader = task02.File{Name: "张三"}
	task02.PrintType(reader)
	var writer task02.Writer = task02.File2{Name: "李四"}
	task02.PrintType(writer)

	//接口实现
	var phone task02.Phone
	phone = new(task02.NokiaPhone)
	phone.Call()
	phone = new(task02.IPhone)
	phone.Call()

}
