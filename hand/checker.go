package hand

import (
	"context"

	"github.com/Barben360/texas-holdem-game/card"
)

type Hander interface {
	// Check checks cards to find what hand they represent
	// there must be exactly 5 cards or ErrBadRequest shall be returned
	Check(ctx context.Context, cards card.Cards5) (Hand, error)
	// GetBest gets the indexes of the best input hands. If there are several, they are equal.
	// There must be at most 10 hands or ErrBadRequest shall be returned
	GetBest(ctx context.Context, hands Hands) ([]int, error)
}
