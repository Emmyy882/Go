//A simple Game program that allow users guess the secret number
package main

import "fmt"

const Secret_Num = 15

var guess_Num int
var i int = 4

func guess_Prompt() {

	fmt.Print("Guess The secret Number: ")
	fmt.Scan(&guess_Num)

}

func main() {

	fmt.Println(">>>>>>>>>WELCOME TO MY GUESS SECRET NUMBER GAME<<<<<<<<")
	fmt.Println(">>>>>>>>>Secret Number Is Between 0 and 20!<<<<<<<<<<<<")

	for ; i > 0; i-- {

		fmt.Println("\nYou Have ", i, " Chances Left!")
		guess_Prompt()

		if guess_Num == Secret_Num {

			fmt.Println("\nHurray, You Won!")
			break

		} else if guess_Num >= 0 && guess_Num <= 10 {

			fmt.Println("\nWrong Guess!\n Tip: Secret Number is above 10 ")
			guess_Prompt()

		} else if guess_Num > 10 && guess_Num != 15 && guess_Num <= 20 {

			fmt.Println("\nYou're Getting Close, Try Again!")
			guess_Prompt()

		} else {

			fmt.Println("\nNumber Not In Range of Secret Number, Try Again!")
			guess_Prompt()

		}
	}

}
