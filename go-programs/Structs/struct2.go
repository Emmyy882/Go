/* more practice on structs using method */
package main

import "fmt"

/* Person struct */
type Person struct {
	name       string
	age        int
	height     string
	complexion string
}

/* info method assigns values to the struct elements and return the values */
func (per *Person) info() (string, int, string, string) {
	fmt.Println("Enter Person Details")

	fmt.Printf("Enter name of person: ")
	fmt.Scanln(&per.name)

	fmt.Printf("\nEnter age of person: ")
	fmt.Scanln(&per.age)

	fmt.Printf("\nEnter height of person in feet(ft): ")
	fmt.Scanln(&per.height)

	fmt.Printf("\nEnter complexion of person: ")
	fmt.Scanln(&per.complexion)

	return per.name, per.age, per.height, per.complexion
}

func main() {
	person1 := Person{} /* initializing an object for the Person struct */

	person1.info()       /* referencing the method to access structs data */
	fmt.Println(person1) /* printing the struct data */

}
