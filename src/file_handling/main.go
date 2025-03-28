package main

import (
	"fmt"
	"os"
)

func main() {

	// ? create a file

	// file, err := os.Create("example.txt")

	// if err != nil {
	// 	fmt.Println("Error while creating file ", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Humpty dumpty sat on a wall "

	// bytes, errors := io.WriteString(file, content+"\nHello world")

	// if errors != nil {
	// 	fmt.Println("Error while creating file ", errors)
	// 	return
	// }
	// fmt.Println("Bytes : ", bytes)

	// fmt.Println("Successfully created a file")

	//? reading a file

	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	fmt.Println("Error while reading file ", err)
	// 	return
	// }
	// defer file.Close()

	// //? create a buffer to read a file content

	// buffer := make([]byte,1024)

	// for{
	// 	n,errors := file.Read(buffer)

	// 	if errors == io.EOF {
	// 		break
	// 	}

	// 	if errors != nil {
	// 		fmt.Println("Error while reading file ", err)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))

	// }

	// ! reading file using io/util | os.ReadFile

	content, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("error reading file ", err)
		return
	}

	fmt.Println(string(content))

}
