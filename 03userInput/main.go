package main

import (
	"bufio"
	"fmt"
	"os"
)

// the below input method is not suitble as of now bcoz it will take only string inputbut we need to different inpouts as use
func main() {
	//to read input from user we need "bufio" package
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Your RollNo : ")

	//comma ok syntax / comma error syntax
	//it is like try catch that if any error comes while input then it will store error in err
	input, err := reader.ReadString('\n') //it means when we hit enter or enter to new line then it will stop takign input
	fmt.Println("My RollNo is : ", input)
	fmt.Println("any error : ", err)
	fmt.Printf("type of input is : %T\n", input) //as of now it will show string

	//we can use this also inplace of err
	input2, _ := reader.ReadString('\n')
	fmt.Println("Enter anyThing")
	fmt.Println(input2)

	//using user input in for loop
	loop()
}

func loop() {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	}
}
