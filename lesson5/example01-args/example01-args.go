package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("num: %d; value: %s\n", i, arg)
	}
}
