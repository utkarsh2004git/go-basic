package main

import "fmt"

func main() {

	var values [5]int
	fmt.Println(values)

	values[0] = 1; 
	values[1] = 2; 
	values[2] = 3; 

	fmt.Println(values)

	for _,val := range values{
		fmt.Printf("The value is : %d\n",val)
	}

	names := [3]string{"Hello","world","hey"}

	fmt.Println(names)
	fmt.Printf("%q",names)

}
