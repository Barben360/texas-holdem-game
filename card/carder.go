package card

import "context"

type Carder interface {
	// Shuffle shuffles cards
	Shuffle52(ctx context.Context, cards *Cards52) error
	// Sort sorts 2 cards by rank first then by suit
	Sort(ctx context.Context, cards Cards, descendent bool) error
}
