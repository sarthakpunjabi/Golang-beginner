
package main

import (
	"fmt"
	"sync"
	
)

func Merge(data1 []float64, data2 []float64) (result []float64) {
	result = make([]float64, len(data1)+len(data2))
	left, right := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case left >= len(data1):
			result[i] = data2[right]
			right++
		case right >= len(data2):
			result[i] = data1[left]
			left++
		case data1[left] < data2[right]:
			result[i] = data1[left]
			left++
		default:
			result[i] = data2[right]
			right++
		}
	}
	return
}

func SingleMergeSort(data []float64) []float64 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}

func MultiMergeSort(data []float64) []float64 {
	if len(data) < 2 {
		return data
	}

	middle := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var data1 []float64
	var data2 []float64

	go func() {
		data1 = MultiMergeSort(data[:middle])
		wg.Done()
	}()

	go func() {
		data2 = MultiMergeSort(data[middle:])

		wg.Done()

	}()

	wg.Wait()
	return Merge(data1, data2)
}

func gomultiplemerge(data []float64) []float64 {

	return MultiMergeSort(data)
}
func main() {

	//static input
	arr := []float64{10,15,20,23,26,50,3,2,6}
	ans := gomultiplemerge(arr)
	fmt.Println(ans)
	
}