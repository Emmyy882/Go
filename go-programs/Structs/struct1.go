//structs are user defined datatypes used to store information of related values
package main

import "fmt"

//declaring a struct
type Person struct {
	name       string
	age        int
	complexion string
	gender     string
}

func main() {

	//creating a new object of type Person
	person1 := Person{
		name:       "Mike Davis",
		age:        24,
		complexion: "light-skinned",
		gender:     "Male",
	}

	fmt.Println("Person Details : ", person1) //printing the person1 details
}
