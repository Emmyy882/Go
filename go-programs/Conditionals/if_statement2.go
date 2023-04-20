//program that uses if and else condition to grant user Entry
package main

import "fmt"

func main() {

	var age int

	fmt.Print("Enter Age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You're In")

	} else {
		fmt.Println("Sorry, You're not old enough for this")
	}

}
