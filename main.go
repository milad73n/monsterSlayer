package main

import "fmt"
import "github.com/milad73n/monsterSlayer/interaction"

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""
	fmt.Println(winner)
	for winner == "" {
		executeGame()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeGame() string {
	currentRound++
	isSpecialRound := currentRound % 3
}

func endGame() {}
