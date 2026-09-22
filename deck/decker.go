package deck

import (
	"context"
)

type Decker interface {
	// Create new deck
	New(ctx context.Context) (Deck, error)
	// Shuffle deck
	Shuffle(ctx context.Context, deck Deck) error
}
