package card

import "context"

type Carder interface {
	// Shuffle shuffles cards
	Shuffle52(ctx context.Context, cards Cards52) error
	// Sort5 sorts cards by rank first then by suit
	Sort5(ctx context.Context, cards Cards5, descendent bool) error
}
