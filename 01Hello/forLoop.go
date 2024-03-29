package main

import "fmt"

func main() {
	fmt.Println("Hello Priyanshu This Side")
	fmt.Println("what about you")

	//for loops in golang
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//we can do like this also
	var j int = 5
	for ; j >= 0; j-- {
		fmt.Println(j)
	}
}
