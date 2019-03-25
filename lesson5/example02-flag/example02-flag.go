package main

import (
	"flag"
	"fmt"
)

func main() {
	fstr := flag.String("fstr", "", "a string")
	fint := flag.Int("fint", 0, "an int")
	flag.Parse()

	fmt.Printf("name: %s; value: %s\n", "fstr", *fstr)
	fmt.Printf("name: %s; value: %d\n", "fint", *fint)
}