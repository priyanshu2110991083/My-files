package main

import "fmt"

func main() {		//it is showing red line but code is running fine
	for i := 0; i <= 5; i++ {
		if i == 4 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
