package main

import (
	"fmt"
	"unsafe"

	"github.com/seerx/base"
)

type bb struct {
	A int
}

type aa struct {
	S string
	C int
	B *bb
}

func main() {
	a := aa{
		S: "123",
		C: 1000,
		B: &bb{
			A: 10,
		},
	}
	base.MemSet(unsafe.Pointer(&a), 0, unsafe.Sizeof(a))

	fmt.Println(a.S)
}
