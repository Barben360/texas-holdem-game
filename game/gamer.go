package game

import (
	"context"

	"github.com/google/uuid"
)

type Gamer interface {
	// New creates a new game
	New(ctx context.Context, nPlayers int8, nTokens Tokens, smallBlind Tokens, bigBlind Tokens) (Game, error)
	// Reset resets a game object. Like New, but reusing all sub-slices
	Reset(ctx context.Context, g *Game, nPlayers int8, nTokens Tokens, smallBlind Tokens, bigBlind Tokens) error
	// Play triggers an action from a player. Game is returned as a value so when the game is over and its data
	// possibly has been reassigned, it is still possible to know who won and when. All sub-slices or pointers are not
	// ensured to be correct in the end as they can be in use in another game, so please only access non-pointer values.
	Play(ctx context.Context, g *Game, playerId uuid.UUID, action Action) (Game, error)
}
