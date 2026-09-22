package deck

import "github.com/Barben360/texas-holdem-game/card"

// Deck represents a 52 cards deck
type Deck struct {
	Cards card.Cards52
}

// Decks represent several decks
type Decks []Deck
