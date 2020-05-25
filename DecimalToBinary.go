package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"math"
)

func decimalToBinary(decimal int) (int, error) {
	// Exception case
	if decimal == 0 {
		return 0, fmt.Errorf("can't devide by: %d", decimal)
	}

	binary := 0
	for k := 0; decimal > 0; {
		binary += decimal%2 * int(math.Pow10(k))
		decimal /= 2
		k++
	}

	return binary, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("You must enter an argument!")
	}

	// Argument is a number
	arg := os.Args[1]
	decimal, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatalf("Argument must be a number: %v\n", err)
	}

	// MAGIC
	binary, err := decimalToBinary(decimal)
	if err != nil {
		log.Fatalf("Error converting decimal to binary: %v", err)
	}

	fmt.Println(binary)
}