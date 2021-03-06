package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"math"
)

func decimalToBinary(decimal int) (int, error) {
	// Negative number
	if decimal < 0 {
		return 0, fmt.Errorf("can't convert negative number: %d", decimal)
	}

	// Number too big
	if decimal > 524287 {
		return 0, fmt.Errorf("number has to be smaller than 524288: %d", decimal)
	}

	// Zero case
	if decimal == 0 {
		return 0, nil
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

	// Argument is negative
	binary, err := decimalToBinary(decimal)
	if err != nil {
		log.Fatalf("Error converting decimal to binary: %v", err)
	}

	// MAGIC
	fmt.Println(binary)
}