package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Hello world");
	// fmt.Println("Please enter something");

	// reader := bufio.NewReader(os.Stdin);

	// input,_ := reader.ReadString('\n');

	// fmt.Println("Thanks for input :",input)

	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64);

	// if(err!= nil){
	// 	fmt.Println(err);
	// } else{
	// 	fmt.Println("Added 1 to your input",numRating+1)
	// }

	fmt.Println("Enter something integer  :")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	intInput,err :=  strconv.ParseFloat(strings.TrimSpace(input),64) ;

	if(err!=nil){
		fmt.Println("Something went wrong",err);
	}else{
		fmt.Println("Here is your added input : ",intInput+10)
	}

}