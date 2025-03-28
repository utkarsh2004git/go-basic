package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of program ")
	data := add(1, 2)
	fmt.Println("data : ", data)

	defer fmt.Println("Middle of program ")
	defer fmt.Println("Ending of program ")

}

//defer works in LIFO
