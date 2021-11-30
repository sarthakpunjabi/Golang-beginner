package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println(" lines of output A")
	done <- true
	fmt.Println(" lines of output B")
	done <- true
	fmt.Println(" lines of output C")
	done <- true
	fmt.Println(" lines of output D")
	done <- true
	fmt.Println(" lines of output E")
	done <- true
	fmt.Println(" lines of output F")
}
func main() {
	done := make(chan bool, 2)
	go worker(done)
	<-done
	time.Sleep(time.Second * 3)
	fmt.Println(" lines of output X")
}
