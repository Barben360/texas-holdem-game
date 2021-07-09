package dft

import (
	"context"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/deck"
	"github.com/Barben360/texas-holdem-game/errors"
	"github.com/sirupsen/logrus"
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

// Create new deck
func (d *Default) New(ctx context.Context, typ deck.Type) (deck.Deck, error) {
	ret := deck.Deck{
		Type: typ,
	}
	switch typ {
	case deck.Type_52:
		cards := make(card.Cards, 52)
		k := 0
		for i := card.CardSuit_Club; i <= card.CardSuit_Spades; i++ {
			for j := card.CardRank_2; j <= card.CardRank_Ace; j++ {
				cards[k] = card.Card{
					Suit: i,
					Rank: j,
				}
				k++
			}
		}
		ret.Cards = cards
	default:
		err := errors.Errorf(errors.ErrUnimplemented, "only 52 cards deck is implemented for now")
		logrus.WithError(err).Error("could not create new deck")
		return ret, err
	}
	return ret, nil
}

// Shuffle deck
func (d *Default) Shuffle(ctx context.Context, deck deck.Deck) error {
	return d.carder.Shuffle(ctx, deck.Cards)
}
