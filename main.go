package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	guessingGame "src/guessingGame/src"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's your name?")
	playerName, _ := reader.ReadString('\n')
	fmt.Println("Hello" + playerName)
	fmt.Println("Letz play a game 😏")

	g := guessingGame.NewGame(playerName, 100, 7)

	for g.GameIsRunning {
		number, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("shit")
		}
		fmt.Printf("NUMBER : %v \n", number)
		intNumber, _ := strconv.Atoi(number)

		fmt.Print(intNumber)
		answer, _ := g.Player.GuessNumber(intNumber, g.Opponent)

		fmt.Print(answer)

		g.CheckGame()
	}
}
