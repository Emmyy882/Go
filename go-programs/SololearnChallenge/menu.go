/* TASK:
You are making a digital Menu App.
The menu choices are stored in an array called menu.
You need to take a number as input, output the corresponding item from the menu.
In case the index is not valid, your app should output "Invalid choice".
The input should indicate the index of the array, meaning 0 corresponds to the first item.
*/

package main

import "fmt"

func main() {

	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	var num int //stores user input
	i := 0

	fmt.Print("Enter number of item choice on the Menu: ")
	fmt.Scan(&num)

	if num > 5 {

		fmt.Println("Invalid choice")
	} else {

		for ; i < len(menu); i++ {

			if num == i {
				fmt.Println(menu[i])
			}
		}
	}

}
