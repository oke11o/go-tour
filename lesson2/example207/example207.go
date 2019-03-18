package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Число %v делится нацело на 2 \n", i)
		} else if i%3 == 0 {
			fmt.Printf("Число %v делится нацело на 3 \n", i)
		} else if i%5 == 0 {
			fmt.Printf("Число %v делится нацело на 5 \n", i)
		}
	}
}
