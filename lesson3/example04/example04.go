package main

import "fmt"

func main() {
	adressBook := make(map[string][]int)

	adressBook["Alex"] = []int{89996543210}
	adressBook["Bob"] = []int{89167243812}
	adressBook["Bob"] = append(adressBook["Bob"], 89155243627)

	fmt.Println(adressBook)

	for name, numbers := range adressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
}
