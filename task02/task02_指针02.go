package task02

// 函数定义：接收一个整数切片指针，将每个元素乘以2

// slicePtr  是一个指针，保存的是切片的地址
// *slicePtr  解引用，得到指针指向的切片本身
// (*slicePtr)[i]  访问该切片的第 i 个元素
func DoubleSlice(slicePtr *[]int) {
	for i := range *slicePtr {
		(*slicePtr)[i] *= 2
	}
}
