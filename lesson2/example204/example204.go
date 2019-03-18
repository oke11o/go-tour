package main

import "fmt"

func main() {
    str := "Test String"
    for i, ch := range str {
        fmt.Println("Код", ch, "символ в строке под номером", i)
    }
}