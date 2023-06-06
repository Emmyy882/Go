//pointers are used to store memory address
package main

import "fmt"

func main() {
	var x int = 34

	var y *int = &x //y is a pointing storing the memory address of x

	fmt.Println("value of x = ", x)
	fmt.Println("pointer pointing to memory address of x = ", y)

	//dereferencing the pointer using * operator
	*y = 12
	fmt.Println("using pointer to change value of x = ", x)
}
