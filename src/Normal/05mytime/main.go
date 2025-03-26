package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Its all about time")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Format("02-01-2006 Monday"))

	createdDate := time.Date(2024, time.May, 19, 5, 0, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006"))

	// nano  := time.Now().Month()
	// fmt.Println(nano)
}
