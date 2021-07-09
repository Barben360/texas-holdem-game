package deck

import "github.com/Barben360/texas-holdem-game/card"

// Type is the type of deck
type Type int8

const (
	Type_UNKNOWN Type = 0 // Unknown type (free deck)
	Type_52      Type = 1 // Classic 52 cards game
)

// Deck represents a cards deck
type Deck struct {
	// Type represents deck type
	Type Type
	// Cards are cards of the deck, ordered as they are physically supposed to be
	Cards card.Cards
}

// Decks represent several decks
type Decks []Deck
