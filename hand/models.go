package hand

import "github.com/Barben360/texas-holdem-game/card"

type Hand struct {
	Type  HandType
	Ranks [5]card.CardRank
	Suit  card.CardSuit // Informational, for flushes only
}

type Hands []Hand

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
	// switch h.Type {
	// // All these cases depend only on primary rank (Three-of-a-kind and full house too, there can't be twice the same primary rank)
	// case HandType_StraightFlush, HandType_FourOfAKind, HandType_Flush,
	// 	HandType_Straight, HandType_Pair, HandType_ThreeOfAKind, HandType_FullHouse:
	// 	equal = h.Rank1 == other.Rank1
	// 	if !equal {
	// 		better = h.Rank1 > other.Rank1
	// 	}
	// // Two-pair case depends on primary and secondary rank
	// case HandType_TwoPair:
	// 	if h.Rank1 == other.Rank1 {
	// 		equal = h.Rank2 == other.Rank2
	// 		if !equal {
	// 			better = h.Rank2 > other.Rank2
	// 		}
	// 	} else {
	// 		better = h.Rank1 > other.Rank1
	// 	}
	// // Two royal flushes are always equal
	// default:
	// 	equal = true
	// }
	// Comparing ranks

	return equal, better
}
