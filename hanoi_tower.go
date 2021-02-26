package main

import "fmt"

func main() {
	process(3, "左", "右", "中")
}

func process(N int, from, to, help string) {
	if N == 1 {
		fmt.Println("Move 1 from", from, "to", to)
	} else {
		process(N-1, from, help, to)
		fmt.Println("Move", N, "from", from, "to", to)
		process(N-1, help, to, from)
	}
}
