package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web Requests!")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Something went wrong ", err)
	}
	defer res.Body.Close()

	data, er := ioutil.ReadAll(res.Body)
	if er != nil {
		fmt.Println("Something went wrong ", err)
	}
	fmt.Println(string(data))
}
