//A simple Game program that allow users guess the secret number
package main

import "fmt"

const Secret_Num = 15

var guess_Num int

func guess_Prompt() {

	fmt.Print("Guess The secret Number: ")
	fmt.Scan(&guess_Num)
	check_Num(guess_Num)

}

func check_Num(guess_Num int) {

	switch guess_Num {

	case guess_Num == Secret_Num:
		fmt.Println("\nHurray, You Won!")
		break

	case guess_Num >= 0 && guess_Num <= 10:
		fmt.Println("\nWrong Guess!\n Tip: Secret Number is above 10 ")
		guess_Prompt()

	case guess_Num > 10 && guess_Num != 15 && guess_Num <= 20:
		fmt.Println("\nYou're Getting Close, Try Again!")
		guess_Prompt()

	default:
		fmt.Println("\nNumber Not In Range of Secret Number, Try Again!")

	}
}

func main() {

	var i int
	fmt.Println(">>>>>>>>>WELCOME TO MY GUESS SECRET NUMBER GAME<<<<<<<<")
	fmt.Println(">>>>>>>>>Secret Number Is Between 0 and 20!<<<<<<<<<<<<")

	for i = 4; i > 0; i -= 1 {

		fmt.Println("\nYou Have ", i, " Chances Left!")
		guess_Prompt()

	}

}
