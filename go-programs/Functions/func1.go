//a function is a block of code that perform a task
package main

import "fmt"

//a basic non-return/no-argument type function to output "Hello, World" when called
func greet() {
	fmt.Println("Hello, World")
}

//declaring a function with string arguments
func greetings(name string) {
	fmt.Println("Hello! ", name)
}

//the main function where all program execution starts
func main() {
	greet() //calling the greet() function

	var name string = "Mike"
	greetings(name) //passing variable name as argument to the function
}
