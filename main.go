package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("welcome to GUESS THE NUMBER challenge !")
	var ans int = rand.Intn(10)
	gameFlow(ans, 0, 9)
	//test()
}

func gameFlow(ans int, start int, end int) {
	var number int
	fmt.Print("Guess a number between " + strconv.Itoa(start) + " and " + strconv.Itoa(end) + " : ")
	fmt.Scanln(&number)
	if(number == ans) {
		fmt.Println("Yay! You won the game.")
	} else {
		if(number < ans) {
			gameFlow(ans, number, end)
		} else {
			gameFlow(ans, start, number)
		}
	}
}