package card

import "context"

type Carder interface {
	// Shuffle shuffles cards
	Shuffle(ctx context.Context, cards Cards) error
	// Draw draws cards from top of pile. If not enough cards, draws all the cards and returns ErrResourceExhausted
	Draw(ctx context.Context, quantity int, cards Cards) (drawn Cards, remaining Cards, err error)
	// Sort cards sorts cards by rank first then by suit
	Sort(ctx context.Context, cards Cards, descendent bool) error
	// Less compares two cards by rank
	// If equal is true, then ranks are the same
	// If comp is false, then card1 <= card2, otherwise card1 >= card2. Undefined behavior if equal is true.
	Less(ctx context.Context, card1 Card, card2 Card) (equal bool, comp bool, err error)
}
