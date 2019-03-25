package main

import "os"
import "strings"
import "fmt"

func main() {
	fmt.Println("PATH:", os.Getenv("PATH"))
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
