//This program shows the Go primtive datatypes and the type of values they store
package main

import "fmt"

func main() {

	var num int = 10         //int stores whole numbers
	var name string = "John" //string stores a group of characters
	var pi float32 = 3.14    //float32 or float64 store numbers with fractional parts
	var active bool = false  //bool can either store true or false

	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", pi, pi)
	fmt.Printf("%v, %T\n", active, active)

}
