//break statement is used to end an iteration when a certain condition is met
package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		if i == 8 {
			break //iteration stops when i == 8 i.e the number 8 is not included in the output
		}
		fmt.Println(i)
	}

}
