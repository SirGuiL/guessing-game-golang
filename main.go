package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("A random number between 1 and 100 is generated, and you have to guess it.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := [10]int64{}

	for i := range attempts {
		fmt.Println("What's your guess?")
		scanner.Scan()

		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)

		attemptInteger, err := strconv.ParseInt(attempt, 10, 64)

		if err != nil {
			fmt.Println("Your attempt must be an integer", err)
			return
		}

		switch {
		case attemptInteger < x:
			fmt.Println("your guess was too low")
		case attemptInteger > x:
			fmt.Println("your guess was too high")
		case attemptInteger == x:
			fmt.Printf("\nCongratulations! You hit the nail on the head\nThe number was: %d\n\nHis attempts were: %v", x, attempts[:i])
			return
		}

		attempts[i] = attemptInteger
	}

	fmt.Printf("You couldn't guess correctly, the number was: %d\nYou had 10 attempts\nHis attempts were: %v", x, attempts)
}
