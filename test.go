package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	x := l.Front()
	if x == nil {
		fmt.Println("xxx")
	}
	fmt.Println(x)
}
