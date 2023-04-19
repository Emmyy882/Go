//a program to calculate the area of a circle
package main

import (
	"fmt"
	"math"
)

func main() {

	var rad float32

	fmt.Print("Enter value for radius: ") //prompting user input for radius
	fmt.Scanf("%f", &rad)                 //reading input from keyboard

	area := math.Pi * rad * rad

	fmt.Println("Area = ", area)

}
