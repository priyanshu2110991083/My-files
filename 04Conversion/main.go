package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter rating for pizza : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// fmt.Println(input+1)		THIS WILL GIVE ERROR AS INPUT AS OF TYPE STRING AND WE CANT ADD 1 TO STRING

	//now we need to convert this into float (we can convert it into another datatypes also)

	//BELOW LINE WILL CREATE ERROR (READ BELOW)
	//newNumber, err := strconv.ParseFloat(input, 64) //we need to specify STRING and BIT SIZE OF DATATYPE
	//strconv STANDS FOR STRING CONVERSION

	//NOW WE NEED IF STATEMENT AS IF ERROR COMES THEN WE NEED TO THROUGH IT ELSE WE NEED TO PRINT NEWNUMBER

	//NOTE -> AS OF NOW IT WILL GIVE ERROR AS ON INPUT WE HAVE SE THAT WHEN WE HIT NEW LINE THEN IT WILL ADD "\N" CHAR TO STRING
	//SO AFTER CONVERTING LET SAY INPUT IS 4 IN STRING
	//NOW IT WILL CONVERTED AS "4\N" SO IN THIS WE CANT ADD 1 SO WE NEED TO DO SOME CHANGES IN BELOW CODE
	//we need to remvoe last char i.e. space from string

	// if err != nil {
	// 	fmt.Println(err)
	// } else { //we need to write else like this only
	// 	fmt.Println("after adding 1 to input : ", newNumber+1)
	// }

	//first we will remove last character(space) from string
	newSrting := strings.TrimSpace(input) //strings is package
	newNumber, err := strconv.ParseFloat(newSrting, 64)

	if err != nil {
		fmt.Println(err)
	} else { //we need to write else like this only
		newNumber++ 		//In Go, unlike some other languages like C or C++, there is no pre-increment or post-increment operator. Instead, you would typically use the increment operator (+) directly to increase the value of a variable.
		fmt.Println("After converting and adding 1 : ", newNumber)	//we vcant use (++newNumber) in go lang
	}

}
