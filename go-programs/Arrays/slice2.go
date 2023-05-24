//More practice on slices
//make() function is used to initialize an array or slice
package main

import "fmt"

func main() {
	var num []int = make([]int, 10) //initializing slice using make() function

	num[0] = 2
	num[1] = 3
	num[2] = 9
	num[3] = 12

	fmt.Println(num)

	//initializing another slice with string elements
	var name []string = make([]string, 5)
	name[0] = "John"
	name[3] = "Lucy"
	name[4] = "Mike"

	fmt.Println("Printing out second slice")
	//using a for-range to output slice elements
	for i, value := range name {
		fmt.Println("index ", i, " value = ", value)
	}
}
