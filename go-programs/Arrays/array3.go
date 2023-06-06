package main

import "fmt"

func main() {

	var arr [10]int = [10]int{4, 3, 43, 10, 1, 76, 88, 6, 3, 4}

	for i := 0; i < len(arr); i++ {
		fmt.Println("index ", i, " value: ", arr[i])
	}

	fmt.Println("Printing another type")

	//another example of array declaration and initialization
	arr2 := [12]int{1, 4, 2, 6, 8, 3, 6, 5, 6, 82}

	for j := 0; j < len(arr); j++ {

		fmt.Println("index ", j, " value: ", arr2[j])
	}
	//can't be there the share

}
