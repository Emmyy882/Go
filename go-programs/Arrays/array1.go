//an array is a non-primitive datatype that is used to store multiple values  of the same datatype
//Elements of an array can be accessed using their index number, which starts from index 0 to (length - 1)

package main

import "fmt"

func main() {

	var arr [5]int //array declaration

	fmt.Println("arr = ", arr) //output the default value 0 since no value has been assigned

	//array initialization
	arr[0] = 2
	arr[1] = 5
	arr[2] = 8
	arr[3] = 4
	arr[4] = 10

	fmt.Println("arr = ", arr)               //outputs the entire array
	fmt.Println("array length = ", len(arr)) //output the length of arr

	//printing the value of a particular index using the index number
	fmt.Println("arr[2] = ", arr[2])

}
