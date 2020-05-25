package main

import (
	"fmt"
	"os"
	"strconv"
)

func decimal_to_binary(number int) {

	if number == 0 { // Default case
		fmt.Println("0")
		return
	}

	decimal := number
	binary := []int{}

	for decimal > 0 {
		binary = append( binary, decimal%2 )
		decimal /= 2
	}

	for i:=1; i<len(binary)+1; i++ {
		fmt.Print( binary[len(binary)-i] )
	}
	fmt.Print("\n")
}

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		decimal, err := strconv.Atoi(arg)
		if err == nil { // Argument is a number
			/* MAGIC */
			decimal_to_binary(decimal)
		} else { // Argument is not a number
			fmt.Println("Argument must be a number!")
		}
	} else { // No arguments
		fmt.Println("You must enter an argument!")
	}

}
