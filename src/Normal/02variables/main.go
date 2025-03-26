package main

import (
	"fmt"
)

const LoginToken string= "YES"


func main() {
	var user string = "Utkarsh"
	fmt.Println(user)
	fmt.Printf("variable is of type %T \n",(user) )

	var isLogin bool = true;
	fmt.Println(isLogin)

	var smallVal uint8 = 255;
	var Val int = 255;
	fmt.Println(smallVal)


	fmt.Println(Val)
	var smallFloat float32 = 123.123456789;
	var bigFloat float64= 123.123456789;
	fmt.Println(smallFloat)
	fmt.Println(bigFloat)

	var hello string
	var demoInt int 
	fmt.Println(hello,demoInt)







	// implicit way of decalaring variable

	var website = "hello world "
	fmt.Println(website)


	//no var style walrus operator | short variable operator

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)



	newVar := "hmmmmm"
	fmt.Println(newVar+"asjdhjashdjash","hello")

}