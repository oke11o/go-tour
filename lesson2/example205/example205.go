package main

import "fmt"

func main() {
    str := "Тестовая строка"
    for i, ch := range str {
        fmt.Println("Код", ch, "символ в строке под номером", i)
    }
}