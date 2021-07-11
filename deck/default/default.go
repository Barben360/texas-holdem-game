package dft

import (
	"context"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/deck"
)

// Default is the default decker implementation
type Default struct {
	carder card.Carder
}

// New creates a new instance of default implementation of decker
func New(ctx context.Context, carder card.Carder) (deck.Decker, error) {
	ret := new(Default)
	ret.carder = carder
	return ret, nil
}

// Create new 52 cards deck
func (d *Default) New(ctx context.Context) (deck.Deck, error) {
	ret := deck.Deck{}
	k := 0
	for i := card.CardSuit_Club; i <= card.CardSuit_Spades; i++ {
		for j := card.CardRank_2; j <= card.CardRank_Ace; j++ {
			ret.Cards[k] = card.Card{
				Suit: i,
				Rank: j,
			}
			k++
		}
	}
	return ret, nil
}

// Shuffle deck
func (d *Default) Shuffle(ctx context.Context, deck deck.Deck) error {
	return d.carder.Shuffle52(ctx, deck.Cards)
}
