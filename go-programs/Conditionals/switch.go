//the switch statement is similar to the if statement where an operation is executed when a condition is met
package main

import "fmt"

func main() {

	var day int = 10

	switch day {

	case 1:
		fmt.Println("Today is Monday!")

	case 2:
		fmt.Println("Today is Tuesday!")

	case 3:
		fmt.Println("Today is Wednesday!")

	case 4:
		fmt.Println("Today is Thursday!")

	case 5:
		fmt.Println("Today is Friday!")

	case 6:
		fmt.Println("Today is Saturday, Hurray it's the Weekend!")

	case 7:
		fmt.Println("Today is Sunday, Hurray it's the Weekend!")

	default:
		fmt.Println("Number doesn't indicate a Day of the Week!")
	}

}
