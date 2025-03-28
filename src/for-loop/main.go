package main

import "fmt"

func main() {

	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	counter := 1
	for {
		if counter <= 10 {

			fmt.Println("5 X", counter, "=", 5*counter)
		} else {
			break
		}
		counter++
	}

	numbers := []int{121,23,334,1213,125};

	for index,val := range numbers{
		fmt.Printf("value at index %d is %d \n",index,val)
	}


	data := "Hello world"

	for index,val := range data{
		fmt.Printf("value at index %d is %c \n",index,val)
	}

	fmt.Printf("%s",data)

}
