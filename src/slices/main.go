package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	numbers = append(numbers, 10)

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println(numbers)
	fmt.Println(len(numbers))	
	fmt.Println("-----------------------------------")	





	names:=[]string{"x","y"}  

	fmt.Println("Slice : ",names)
	fmt.Println("Length : ",len(names))
	fmt.Println("Capacity : ",cap(names))
		
	names = append(names, "a")
	fmt.Println("Slice : ",names)
	fmt.Println("Length : ",len(names))
	fmt.Println("Capacity : ",cap(names))

	names = append(names, "a")
	fmt.Println("Slice : ",names)
	fmt.Println("Length : ",len(names))
	fmt.Println("Capacity : ",cap(names))

	names = append(names, "a")
	fmt.Println("Slice : ",names)
	fmt.Println("Length : ",len(names))
	fmt.Println("Capacity : ",cap(names))


	fmt.Println("-----------------------------------")	
	//make function
 
	integers := make([]int,3,5)      //make(type , length , capacity)

	//if size is given the 0 values will stored for that size
	// if size is given as 0 then it will initialize empty slice eg. 
	// nums := make([]int,0)

	integers[0] = 1;
	integers[1] = 2;
	integers[1] = 3;
	fmt.Println("Slice Integers : ",integers)
	fmt.Println("Length : ",len(integers))
	fmt.Println("Capacity : ",cap(integers))

	integers = append(integers, 1)
	integers = append(integers, 1)
	integers = append(integers, 1)
	integers = append(integers, 1)
	fmt.Println("Slice Integers : ",integers)
	fmt.Println("Length : ",len(integers))
	fmt.Println("Capacity : ",cap(integers))

	integers = append(integers, 1)
	integers = append(integers, 1)
	integers = append(integers, 1)
	integers = append(integers, 1)
	fmt.Println("Slice Integers : ",integers)
	fmt.Println("Length : ",len(integers))
	fmt.Println("Capacity : ",cap(integers))

}
