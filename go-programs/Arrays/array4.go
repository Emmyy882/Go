package main

import "fmt"

//function to receive and output elements of array
func print(arr []int) {

	for i := 0; i < len(arr); i++ {

		//prompting user input
		fmt.Print("Enter ", i, "th number element of the array: ")
		fmt.Scan(&arr[i]) //reading input from keyboard
	}
	fmt.Println(arr) //printing the elements of the array
}

func main() {
	//declaring array
	var arr1 []int = make([]int, 5)

	//calling the print function
	print(arr1)
}
