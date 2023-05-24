//Slices are like arrays just that they do not include size in their declaration
package main

import "fmt"

func main() {

	slice1 := []int{2, 3, 8, 4, 5} //initializing slice
	slice1[0] = 5                  //reassigning value at index 0
	slice1[1] = 14                 //reassigning value at index 1
	slice1[2] = 20                 //reassigning value at index 2

	fmt.Println(slice1)
}
