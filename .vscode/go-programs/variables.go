//variables are containers which store values and whose values can also be changed
package main

import "fmt"

func main() {

	var name string = "Mike"
	var age int = 22

	fmt.Println("Name is ", name, " and Age is ", age)

	name = "peter" //re-assignment of the variable 'name'
	age = 18       //re-assignment of the variable 'age'
	fmt.Println("Name is ", name, " and Age is ", age)

}
