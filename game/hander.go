package game

import (
	"context"

	"github.com/google/uuid"
)

type Gamer interface {
	// New creates a new game
	New(ctx context.Context, nPlayers int8, nTokens Tokens, smallBlind Tokens, bigBlind Tokens) (Game, error)
	// Reset resets a game object. Like New, but reusing resources
	Reset(ctx context.Context, g *Game, nPlayers int8, nTokens Tokens, smallBlind Tokens, bigBlind Tokens) error
	// Play triggers an action from a player
	Play(ctx context.Context, g *Game, playerId uuid.UUID, action Action) (GameData, error)
}
