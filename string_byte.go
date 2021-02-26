package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Println(b)
	c := []byte(a)
	fmt.Println(c)
}

//# go build -gcflags "-N -l" string_byte.go
//# gdb string_byte
//(gdb) b 15
//(gdb) info locals
//(gdb) ptype a
//(gdb) ptype b
//(gdb) ptype c
