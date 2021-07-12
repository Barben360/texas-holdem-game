package game

import (
	"time"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/deck"
	"github.com/google/uuid"
)

// Game represents a game
type Game struct {
	Id      uuid.UUID
	Data    GameData    // Game data. Can be safely accessed contrary to other members
	Deck    deck.Deck   // Game deck
	Players Players     // Players
	Table   card.Cards5 // Always five even if not all revealed
}

type GameData struct {
	StartedAt        time.Time // When the game started
	EndedAt          time.Time // When the game ended (time's zero if still running)
	WonBy            int8      // Position of the player who won after game has ended (-1 if game is not over)
	Dealer           int8      // Position of the dealer for current round
	Playing          int8      // Position of player playing
	RemainingPlayers int8      // Number of remaining players in game
	Round            int       // Current round
	GameStep         GameStep  // Current round's game step
	Error            error     // Error details if game is in error. Uses errors.GenericError type
}

type GameStep int8

const (
	GameStep_Error GameStep = 0 // Game is in unrecoverable error (e.g, too many rounds)
	GameStep_End   GameStep = 1 // Game has ended
	GameStep_Flop  GameStep = 2 // 3 cards revealed
	GameStep_Turn  GameStep = 3 // 4 cards revealed
	GameStep_River GameStep = 4 // 5 cards revealed
)

// Player represents a player
type Player struct {
	Id          uuid.UUID
	Hand        card.Cards2 // Hand (always 2 cards, ordered)
	TokensOwned Tokens      // Owned tokens (does not count current bet)
	TokensBet   Tokens      // Bet tokens
	Rank        int8        // If 0, the player is still in game
	HasFolded   bool        // If true, has folded at current round
}

type Players []Player

type Tokens int

type Action struct {
	Type   ActionType // Action type
	Amount Tokens     // Only when playing, not folding
}

type ActionType int8

const (
	ActionType_Play ActionType = 0
	ActionType_Fold ActionType = 1
)
