package main

import "fmt"

func main() {

	var i int = 12 //store whole numbers
	var j float32 = 8.2 //stores fractional parts
	var name string = "Mike" //stores group of characters
	var light bool = false //can either be true or false

	fmt.Printf("%v, %T\n ", i, i)
	fmt.Printf("%v, %T\n ", j, j)
	fmt.Printf("%v, %T\n ", name, name)
	fmt.Printf("%v, %T\n ", light, light)

}
