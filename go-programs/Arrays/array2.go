//more practice on arrays
package main

import "fmt"

func main() {

	var name [4]string

	//using for loop to prompt and assign values from user
	for i := 0; i < len(name); i++ {

		fmt.Print("Enter Name ", i, " : ")
		fmt.Scan(&name[i])
	}
	fmt.Println("name = ", name)

	//reassigning the value in index 2
	name[2] = "Lucy"
	fmt.Println("name = ", name)

}
