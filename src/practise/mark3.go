package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	var b = &arr[1]
	*b = 5
	fmt.Println(arr)

}
