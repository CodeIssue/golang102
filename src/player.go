package guessingGame

import "errors"

type Player struct {
	Name string
}

func (p *Player) GuessNumber(n int, o *Opponent) (string, error) {
	if n > o.Number {
		return "Your n is to big, Bro!\n", nil
	}
	if n < o.Number {
		return "Your n is to small, Yo!\n", nil
	}
	if n == o.Number {
		return "You're Right!\n", nil
	}
	return "Your Number is ERRRR", errors.New("notPossibleButError")
}
