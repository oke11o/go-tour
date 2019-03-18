package main

import "fmt"

func main() {
	ourSlice := []int{98, 93, 77, 82, 83}

	average(ourSlice)
	average([]int{10, 20, 14, 100})
}

func average(values []int) {
	total := 0
	for _, v := range values {
		total += v
	}
	fmt.Println(total / int(len(values)))
}
