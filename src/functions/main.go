package main

import "fmt"

func SimpleFunction() {
	fmt.Println("This is Simple function")
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64 , error) {
	if b == 0 {
		return 0 , fmt.Errorf("Can't divide by zero")
	}
	return a/b ,nil
}

func main() {

	fmt.Println("This is MAIN function")
	SimpleFunction()
	fmt.Println(add(1, 2))

	ans ,err :=divide(7,0)

	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(ans)
	}


	
}
