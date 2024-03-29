package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My Time")
	myTime := time.Now() //this will provide date with other information also but we need only date
	fmt.Println(myTime)

	//to print only date we can do like this
	fmt.Println(myTime.Format("01-02-2006")) //now it will print date only in this format
	//NOTE - WE NEED TO GIVE SAME DATE IN SAME ORDER AS IT IS MENTIONED IN DOCUMENTATION

	//TO PRINT DAY NAME ALONG WITH DATE
	fmt.Println(myTime.Format("01-02-2006 Monday")) //now it will print actual day on which you are running this code

	//see documentation for more details

	//to create our own date
	createdDate := time.Date(2003, time.September, 25, 4, 4, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
