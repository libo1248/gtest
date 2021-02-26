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
	fmt.Println([]byte(a))
}
