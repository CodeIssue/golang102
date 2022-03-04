package guessingGame

import (
	"fmt"
	"math/rand"
)

type Opponent struct {
	Number int
}

func NewOpponent(n uint) *Opponent {
	myNumber := rand.Intn(int(n)) + 1
	fmt.Print(myNumber)
	return &Opponent{
		Number: myNumber,
	}
}
