package main

import "fmt"

func fibonacci() func() int {
	val1, val2 := 0, 1
	return func() int {
		val1, val2 = val2, val1 + val2
		return val2
	}
}

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
