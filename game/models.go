package game

import (
	"time"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/deck"
	"github.com/google/uuid"
)

// Game represents a game
type Game struct {
	Id        uuid.UUID
	StartedAt time.Time // When the game started
	EndedAt   time.Time // When the game ended
	WonBy     int8      // Position of the player who won after game has ended
	Deck      deck.Deck // Game deck
	Players   Players
	Dealer    int8
	Playing   int8       // Position of player playing
	Table     card.Cards // Always five even if not all revealed
}

// Player represents a player
type Player struct {
	Id        uuid.UUID
	Hand      card.Cards
	Tokens    Tokens
	Rank      int8 // If 0, the player is still in game
	HasFolded bool // If true, has folded at current round
}

type Players []Player

type Tokens int

type Action struct {
	Type   ActionType
	Amount Tokens // Only when playing, not folding
}

type ActionType int8

const (
	ActionType_Play = 0
	ActionType_Fold = 1
)
