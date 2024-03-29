package main

import "fmt"

//if we are declaring variable with 1st letter capital then it means we are declaring it as PUBLIC
const Login bool = true

//here "Login" act as public variable as its 1st character is capital

func main() {
	fmt.Println("Variables")
	fmt.Println("Login : ", Login)
	//if we write only this line w/o printing or using it so this will give error
	var username string = "Priyanshu"
	fmt.Println(username)

	fmt.Printf("Its Type is : %T \n", username)
	withoutDataType()
	withoutVar()
	Integer1()
	Boolean1()
	uint1()
	float1()
	default1()
}

func withoutVar() {
	// NOTE -> INSIDE ANY METHOD WE ARE ALLOWED TO INITIALISE VARIABKE LIKE THISBUT GLOBALLY OR OUTSIDE THE FUNCTION WE ARE NOT ALLOWED
	randomNumber := 100000 // ":=" THIS IS KNOWN AS "VOLROUS OPERATOR"
	fmt.Println(randomNumber)
	fmt.Printf("type var is : %T \n", randomNumber)
}

func withoutDataType() {
	//if we are not specifing the datatype of variable then LEXER will autumatically correct it
	var name = "Priyanshu"
	fmt.Println(name)
	fmt.Printf("type : %T \n", name)
}

//integer datatype
func Integer1() {
	var n int = 10
	fmt.Println(n)
	fmt.Printf("Type is : %T \n", n)
}

//boolean datatype
func Boolean1() {
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("the datatype if isLoggedIn is : %T \n", isLoggedIn)

}

//uint datatype
func uint1() {
	var smallVal uint8 = 255 //if we go after 255 then uint8 get out of range
	fmt.Println(smallVal)
	fmt.Printf("the datatype if smallVal is : %T \n", smallVal)

}

//float32 and floaf64
func float1() {
	//float32 provide upto 5 values after decimal
	var smallFloat float32 = 255.2432342342322342
	fmt.Println(smallFloat)
	fmt.Printf("datatype is : %T \n", smallFloat)

	//to get more values after decimal then we can use float64 here
	var bigFloat float64 = 255.2432342342322342
	fmt.Println(bigFloat)
	fmt.Printf("datatype is : %T \n", bigFloat)
}

//default values and alises
func default1() {
	//default value of int (when we had not declared) is 0
	var anotherValue int
	var anotherFloat float32 // 0
	var anotherBool bool     // flase
	var anotherString string //default value is empty string
	fmt.Println(anotherValue)
	fmt.Println(anotherFloat)
	fmt.Println(anotherBool)
	fmt.Println(anotherString)

}
