package main

import "fmt"

//why pointers ?

//suppose we have created one array and stored values / variables init
//so when we are passing such values / variabkes in function then sometimes copy of these variables get converted and go inside function
//but we need to pass reference of that vsriables so that when we do any change in variables so it should reflect inside array
//that's why we need to use pointers

func main() {
	fmt.Println("Pointers")
	//initial value of pointer is "nil"

	var ptr *int // "*" means creating a pointer
	fmt.Println("Default value of pointer is ", ptr)

	myNumber := 10
	//"pointer" is referencing to the address of "myNumber"
	var pointer = &myNumber //"&" means we are referncing from something
	fmt.Println(pointer)    //prints address

	fmt.Println(*pointer) //derefrencing	so it will print 10

	//var a *int = 10		//it is not possible as we are assgning integer value in pointer
	//we can use this
	value := 10;
	var a *int = &value	// "a" is pointing to address of "value" variable
	fmt.Println(a)
	fmt.Println(*a)
	//fmt.Println(a)

}
