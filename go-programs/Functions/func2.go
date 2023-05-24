//more practice on functions using return type functions

package main

import "fmt"

//declaring a return type function with two arguments
func add(x int, y int) int {

	return (x + y)
}

//function to return the difference of two numbers
func sub(x, y int) int {

	return (x - y)
}

//function to return the multiplication of two numbers
func mult(x, y int) int {

	return (x * y)
}

//function to return the division of two numbers
func div(x, y int) int {

	return (x / y)
}

func main() {

	x := 10
	y := 5
	var result int

	result = add(x, y) //assigning the add() function to the result variable to get the reesult of the operation
	fmt.Println("sum = ", result)

	result = sub(x, y)
	fmt.Println("sub = ", result)

	result = mult(x, y)
	fmt.Println("Mult = ", result)

	result = div(x, y)
	fmt.Println("Div = ", result)
}
