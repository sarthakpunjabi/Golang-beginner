package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var result = make(map[string]int)
	arr := strings.Fields(s)
	for _, data := range arr {
		if _, ok := result[data]; ok {
			result[data] += 1
		} else {
			result[data] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
