package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/pbreedt/stdio/input"
)

func main() {
	myNumber := getRandomNumber()

	for tries := 0; tries < 10; tries++ {
		userGuess := getUserGuess()
		if userGuess == myNumber {
			fmt.Println("CORRECT!")
			os.Exit(0)
		} else if userGuess < myNumber {
			fmt.Println("Your guess was LOW, try again")
		} else {
			fmt.Println("Your guess was HIGH, try again")
		}
	}
	fmt.Println("FAILED :(", "My number was", myNumber)
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100) + 1
}

func getUserGuess() int {
	userInt, err := input.ReadInt("Guess my number:")

	if err != nil {
		log.Fatal(err)
	}

	return userInt
}
