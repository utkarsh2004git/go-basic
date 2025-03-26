package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple , orange , banana"

	parts := strings.Split(data, " , ")

	fmt.Println(parts)

	for _,value := range parts{
		fmt.Println("Fruit "+value)
	}

	str := "one two three tows four"

	fmt.Println("Count of two is ",strings.Count(str,"two"))


	str2 := "      helloo    "

	trimmed := strings.TrimSpace(str2);

	fmt.Println(trimmed)

	var newStr string = "my name is utkarsh"
	
	fmt.Println(strings.ToTitle(newStr))

	result := strings.Map( func(r rune) rune {
		if(r =='a'){
			return '$'
		}
		return r
	},newStr)

	fmt.Println(result)
	
	a1 := "Utkarsh"
	a2 := "Shahare"	
	a3 := "Company"
	name:= strings.Join([]string{a1,a2,a3}," | ")
	
	fmt.Println(name)

}
