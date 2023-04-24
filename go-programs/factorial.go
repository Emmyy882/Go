//program to calculate the factorial of a number
package main

import "fmt"

var num int
var fact int = 1

func factorial(num int) {

	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println(num, " Factorial = ", fact)
}

func main() {

	fmt.Print("Enter Number To Calculate Factorial: ")
	fmt.Scan(&num)

	factorial(num)
}
