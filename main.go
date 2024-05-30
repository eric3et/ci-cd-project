package main

import (
	"fmt"
)

func main() {

	input := []float64{1, 2, 3}

	fmt.Println(average(input))
}

func average(arr []float64) float64 {
	var result float64 = 0.0
	n := float64(len(arr))
	for i := range arr {
		result += float64(arr[i])
	}

	return result / n
}
