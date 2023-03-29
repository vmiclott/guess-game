package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max := 100
	solution := rand.Intn(max)
	score := max
	lives := 2
	var guess int

	fmt.Println("Welcome")
	fmt.Printf("Guess a number in the interval [0, %d]\n", max-1)
	fmt.Println("If you guess correctly, you win")
	fmt.Println("If you guess incorrectly, you are told whether you guessed too low or too high")
	fmt.Println("You start with 100 points and 2 lives")
	fmt.Println("Every time you guess too high, you lose a life and a point")
	fmt.Println("Every time you guess too low, you lose a point")
	fmt.Println("If you run out of lives, you lose")
	for lives > 0 {
		fmt.Print("Make a guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			panic("Failed to read input")
		}
		if guess >= max || guess < 0 {
			fmt.Printf("Your guess is not in the interval [0, %d]\n", max-1)
			continue
		}
		if guess < solution {
			score--
			fmt.Println("Your guess is too low, try again")
			continue
		}
		if guess > solution {
			score--
			lives--
			if lives == 0 {
				fmt.Println("Your guess is too high, no lives remaining")
				fmt.Println("You lose")
				os.Exit(1)
			}
			fmt.Println("Your guess is too high, 1 life remaining, try again")
			continue
		}
		fmt.Printf("You win, with a score of %d\n", score)
		break
	}
}
