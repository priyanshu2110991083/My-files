package main

import "fmt"

func main() {
	value := 10
	var ptr *int = &value

	fmt.Println("initial value ", *ptr)

	*ptr = *ptr * 2    //now here 1st we derefering ptr so it will convert to 10 and then we are multiplying it with 2 so it become 20 and then at last we are storing 20 on address where ptr is pointing
	fmt.Println(value) //so value of "value" variable get changed
	fmt.Println(*ptr)  //its value is also changed
}
