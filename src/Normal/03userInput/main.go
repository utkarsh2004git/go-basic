package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome Here"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter anything : ")

	// comma ok || comma error syntax 

	input , _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("%T",input)

}