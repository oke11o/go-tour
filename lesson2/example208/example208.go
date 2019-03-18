package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Ноль")
		case 1:
			fmt.Println("Один")
		case 2:
			fmt.Println("Два")
		case 3:
			fmt.Println("Три")
		case 4:
			fmt.Println("Четыре")
		case 5:
			fmt.Println("Пять")
		default:
			fmt.Println("Неизвестный номер")
		}
	}
}
