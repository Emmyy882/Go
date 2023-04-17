//a program to calculate arithmetic operations based on user inputs

package main

import "fmt"

func main() {

	var firstNum, secondNum int
	var operator string

	fmt.Println("Enter first Number for calculation: ")
	fmt.Scanf("%d", &firstNum) //reading input from keyboard

	fmt.Println("Enter operator Number for calculation: ")
	fmt.Scanf("%s", &operator)

	fmt.Println("Enter second Number for calculation: ")
	fmt.Scanf("%d", &secondNum)

	if operator == "+" {

		sum := firstNum + secondNum
		fmt.Println("Result of ", firstNum, " + ", secondNum, " = ", sum)

	} else if operator == "-" {

		sub := firstNum - secondNum
		fmt.Println("Result of ", firstNum, " - ", secondNum, " = ", sub)

	} else if operator == "*" {

		mult := firstNum * secondNum
		fmt.Println("Result of ", firstNum, " * ", secondNum, " = ", mult)

	} else if operator == "/" {

		div := firstNum / secondNum
		fmt.Println("Result of ", firstNum, " / ", secondNum, " = ", div)

	} else {

		fmt.Println("Operator mismatched")
	}

}
