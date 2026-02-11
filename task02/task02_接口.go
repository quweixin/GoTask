package task02

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReaderWriter interface {
	Reader
	Writer
}
type File struct {
	Name string
}
type File2 struct {
	Name string
}

func (f File) Read() string {
	return "Reading from file"
}

func (f File) Write(data string) {
	fmt.Println("Writing to file:", data)
}
func (f2 File2) Write(data string) {
	fmt.Println("Writing to file 2:", data)
}

func PrintType(val interface{}) {
	switch v := val.(type) {
	case Reader:
		fmt.Println("Reader:", v.Read())
	case Writer:
		v.Write("Hello, World!")
	case ReaderWriter:
		fmt.Println("ReaderWriter:", v.Read())
		v.Write("Hello, World!")
	default:
		fmt.Println("Unknown type")
	}
}

type Phone interface {
	Call()
}

type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct{}

func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}
