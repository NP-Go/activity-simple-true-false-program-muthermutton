package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88

	//Insert your code from here
	if value == secretValue {
		resultMessge = "Well Done! Your guess is correct"
	} else if value < secretValue {
		resultMessge = "Too low, try again next time!"
	} else {
		resultMessge = "Too high, try again next time!"
	}
	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	compare(guess)
}
