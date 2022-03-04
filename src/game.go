package guessingGame

type Game struct {
	Player        *Player
	Opponent      *Opponent
	Round         uint
	MaxRounds     uint
	GameIsRunning bool
}

// create new game
// NewGame(playerName, maxNumber, maxRounds) *Game
// playerName : Name of the Player
// maxNumber : Number between 0 and maxNumber
// maxRounds : max Rounds of Guessing
func NewGame(playerName string, maxNumber uint, maxRounds uint) *Game {
	return &Game{
		Player:        &Player{playerName},
		Opponent:      NewOpponent(maxNumber),
		Round:         1,
		MaxRounds:     maxRounds,
		GameIsRunning: true,
	}
}

func (g *Game) CheckGame() {
	g.Round += 1

	if g.Round >= g.MaxRounds {
		g.GameIsRunning = false
	}
}
