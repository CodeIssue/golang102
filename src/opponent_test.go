package guessingGame

import (
	"log"
	"testing"
)

func TestNewOpponent(t *testing.T) {
	o := NewOpponent(100)

	log.Print(o)
}
