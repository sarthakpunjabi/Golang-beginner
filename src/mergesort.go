package main
//21183147

import (
	"fmt"
	
)

func Merge(data1 []float64, data2 []float64) (result []float64) {
	leftone, rightone := 0, 0
	result = make([]float64, len(data1)+len(data2))

	for i := 0; i < cap(result); i++ {
		switch {
		case leftone >= len(data1):
			result[i] = data2[rightone]
			rightone++
		case rightone >= len(data2):
			result[i] = data1[leftone]
			leftone++
		case data1[leftone] < data2[rightone]:
			result[i] = data1[leftone]
			leftone++
		default:
			result[i] = data2[rightone]
			rightone++
		}
	}
	return
}

func MultiMergeSort(data []float64, res chan []float64) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []float64)
	rightChan := make(chan []float64)
	middle := len(data) / 2

	go MultiMergeSort(data[:middle], leftChan)
	go MultiMergeSort(data[middle:], rightChan)

	data1 := <-leftChan
	data2 := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- Merge(data1, data2)
	return
}

func RunMultiMergeSort(data []float64) (multiResult []float64) {
	res := make(chan []float64)
	go MultiMergeSort(data, res)
	multiResult = <-res
	return
}


func main() {

	//static input
	arr := []float64{10,15,20,23,26,50,3,2,6}
	ans := RunMultiMergeSort(arr)
	fmt.Println(ans)

}