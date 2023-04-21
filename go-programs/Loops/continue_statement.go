//program that shows how the continue statement works
package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			continue //ignores numbers divisible by 2; continue statement is used to skip an operation when a certain conditions is met
		}
		fmt.Println(i)
	}
}
