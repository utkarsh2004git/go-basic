package main

import (
	"fmt"
	"time"
)

func main() {
	//yyyy-mm-dd
	//2006-01-02 15:04:05
	// 2 Jan 2006 , 3:04:05 PM

	currentTime := time.Now().Local()

	fmt.Println("current ", currentTime)

	formattedDate := currentTime.Format("02-01-2006, Mon 03:04:05 PM")
	fmt.Println(formattedDate)

	dateStr := "2004-05-19"

	newFormatTime, _ := time.Parse("2006-01-02", dateStr)
	fmt.Println(newFormatTime)
	fmt.Println(newFormatTime.Format("02-01-2006"))

	new_date := currentTime.Add(50 * time.Minute)

	fmt.Println("current ", currentTime)
	fmt.Println(new_date)

}
