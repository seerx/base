package base

import "unsafe"

// MemCopy 可以拷贝数组、结构体
func MemCopy(dest unsafe.Pointer, src unsafe.Pointer, n uintptr) {
	destPtr := uintptr(dest)
	srcPtr := uintptr(src)
	var i uintptr
	for i = 0; i < n; i++ {
		destByte := (*byte)(unsafe.Pointer(destPtr + i))
		*destByte = *(*byte)(unsafe.Pointer(srcPtr + i))
	}
}

// MemSet 可以填充数组、结构体
func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}

// BZero 清零，结构体或数组
// var arr [10]int32
// MemSet(unsafe.Pointer(&arr), 0, unsafe.Sizeof(arr))
// fmt.Printf("%+v\n", arr)
func BZero(s unsafe.Pointer, n uintptr) {
	MemSet(s, 0, n)
}
