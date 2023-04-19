//program to calculate simple interest
package main

import "fmt"

func main() {

	var principal, rate, time float32

	fmt.Print("Enter amount: ")
	fmt.Scan(&principal)

	fmt.Print("Enter rate: ")
	fmt.Scan(&rate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&time)

	interest := (principal * rate * time) / 100
	fmt.Println("Simple Interest = ", interest)

}
