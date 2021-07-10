package game

import (
	"time"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/deck"
	"github.com/google/uuid"
)

// Game represents a game
type Game struct {
	Id      uuid.UUID
	Data    GameData   // Game data. Can be safely accessed contrary to other members
	Deck    deck.Deck  // Game deck
	Players Players    // Players
	Table   card.Cards // Always five even if not all revealed
}

type GameData struct {
	StartedAt        time.Time // When the game started
	EndedAt          time.Time // When the game ended (time's zero if still running)
	WonBy            int8      // Position of the player who won after game has ended (-1 if game is not over)
	Dealer           int8      // Position of the dealer for current round
	Playing          int8      // Position of player playing
	RemainingPlayers int8      // Number of remaining players in game
	Round            int       // Current round
	GameStep         GameStep  // Current round's game step
	Error            error     // Error details if game is in error. Uses errors.GenericError type
}

type GameStep int8

const (
	GameStep_Error GameStep = 0 // Game is in unrecoverable error (e.g, too many rounds)
	GameStep_End   GameStep = 1 // Game has ended
	GameStep_Flop  GameStep = 2 // 3 cards revealed
	GameStep_Turn  GameStep = 3 // 4 cards revealed
	GameStep_River GameStep = 4 // 5 cards revealed
)

// Player represents a player
type Player struct {
	Id          uuid.UUID
	Hand        card.Cards // Hand (always 2 cards, ordered)
	TokensOwned Tokens     // Owned tokens (does not count current bet)
	TokensBet   Tokens     // Bet tokens
	Rank        int8       // If 0, the player is still in game
	HasFolded   bool       // If true, has folded at current round
}

type Players []Player

type Tokens int

type Action struct {
	Type   ActionType
	Amount Tokens // Only when playing, not folding
}

type ActionType int8

const (
	ActionType_Play ActionType = 0
	ActionType_Fold ActionType = 1
)

type Hand struct {
	Type          HandType
	PrimaryRank   card.CardRank
	SecondaryRank card.CardRank
	Suit          card.CardSuit // Informational, for flushes only
}

// HandType represents hand type, ordered by ranking
type HandType int8

const (
	HandType_SingleCard    HandType = 0
	HandType_Pair          HandType = 1
	HandType_TwoPair       HandType = 2
	HandType_ThreeOfAKind  HandType = 3
	HandType_Straight      HandType = 4
	HandType_Flush         HandType = 5
	HandType_FullHouse     HandType = 6
	HandType_FourOfAKind   HandType = 7
	HandType_StraightFlush HandType = 8
	HandType_RoyalFlush    HandType = 9
)

func (h Hand) IsBetterThan(other Hand) (equal, better bool) {
	if h.Type < other.Type {
		return false, false
	}
	if h.Type > other.Type {
		return false, true
	}
	// Same type
	switch h.Type {
	// All these cases depend only on primary rank (Three-of-a-kind and full house too, there can't be twice the same primary rank)
	case HandType_StraightFlush, HandType_FourOfAKind, HandType_Flush,
		HandType_Straight, HandType_Pair, HandType_ThreeOfAKind, HandType_FullHouse:
		equal = h.PrimaryRank == other.PrimaryRank
		if !equal {
			better = h.PrimaryRank > other.PrimaryRank
		}

	// Two-pair case depends on primary and secondary rank
	case HandType_TwoPair:
		if h.PrimaryRank == other.PrimaryRank {
			equal = h.SecondaryRank == other.SecondaryRank
			if !equal {
				better = h.SecondaryRank > other.SecondaryRank
			}
		} else {
			better = h.PrimaryRank > other.PrimaryRank
		}
	// Two royal flushes are always equal
	default:
		equal = true
	}
	return equal, better
}
