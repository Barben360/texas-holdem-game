package dft

import (
	"context"
	"sync"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/errors"
	"github.com/Barben360/texas-holdem-game/hand"
	"github.com/Barben360/texas-holdem-game/services"
	"github.com/sirupsen/logrus"
)

type Default struct {
	svcs         services.Services
	carder       card.Carder
	ranksMapPool sync.Pool
}

func New(ctx context.Context, svcs services.Services, carder card.Carder) (hand.Hander, error) {
	ret := new(Default)
	ret.svcs = svcs
	ret.carder = carder
	ret.ranksMapPool = sync.Pool{
		New: func() interface{} {
			return make(map[card.CardRank]int, 5)
		},
	}
	return ret, nil
}

// Check checks cards to find what hand they represent
// there must be exactly 5 cards or ErrBadRequest shall be returned
func (d *Default) Check(ctx context.Context, cards card.Cards5) (hand.Hand, error) {
	panic("not implemented") // TODO: Implement (call workers)
}

func (d *Default) check(ctx context.Context, cards card.Cards5) (hand.Hand, error) {
	ret := hand.Hand{
		Type: hand.HandType_SingleCard,
	}

	// Sorting cards
	// here cards object is a copy as it is a const size slice
	err := d.carder.Sort(ctx, cards[:], false)
	if err != nil {
		errMsg := "could not check cards"
		logrus.WithError(err).
			Error(errMsg)
		return ret, errors.Errorf(errors.ErrBadRequest, errMsg)
	}

	// Checking hand
	firstCardSuit := cards[0].Suit
	sameSuit := true
	// Checking if same suit
	for i := 1; i < 5; i++ {
		if cards[i].Suit != firstCardSuit {
			sameSuit = false
			break
		}
	}
	isStraight := true
	for i := 1; i < 5; i++ {
		if cards[i].Rank != cards[i-1].Rank+card.CardRank(1) {
			isStraight = false
			break
		}
	}
	// Mapping cards rank to the number of times they appear
	rankCount := d.ranksMapPool.Get().(map[card.CardRank]int)
	defer func() {
		// Cleaning map
		for key := range rankCount {
			delete(rankCount, key)
		}
		d.ranksMapPool.Put(rankCount)
	}()
	for _, c := range cards {
		rankCount[c.Rank]++
	}
	if len(rankCount) == 5 {
		// Ranks will be forcefully ordered by decreasing order (includes all flushes, straight and single card)
		for i := 0; i < 5; i++ {
			ret.Ranks[i] = cards[4-i].Rank
		}
	}

	// If same suit, this is a flush
	if sameSuit {
		ret.Suit = firstCardSuit
		if isStraight {
			if ret.Ranks[0] == card.CardRank_Ace {
				// This is a royal flush
				ret.Type = hand.HandType_RoyalFlush
				return ret, nil
			}
			// This is a classical straight flush
			ret.Type = hand.HandType_StraightFlush
			return ret, nil
		}
		// This is a simple flush
		ret.Type = hand.HandType_Flush
		return ret, nil
	}

	// If we reach here and isStraight is true, then this is a simple straight
	if isStraight {
		ret.Type = hand.HandType_Straight
		return ret, nil
	}

	if len(rankCount) == 2 {
		// Only 2 ranks on 5 cards, this is a four-of-a-kind or a full house
		// Primary rank may not be the one calculated before
		for r, val := range rankCount {
			if val == 4 {
				ret.Type = hand.HandType_FourOfAKind
				ret.Ranks[0] = r
			} else if val == 3 {
				ret.Type = hand.HandType_FullHouse
				ret.Ranks[0] = r
			} else {
				ret.Ranks[1] = r
			}
		}
		return ret, nil
	}
	if len(rankCount) == 3 {
		// Only 3 ranks on 5 cards, this is a three-of-a-kind or a two pair
		for r, val := range rankCount {
			if val == 3 {
				ret.Type = hand.HandType_ThreeOfAKind
				ret.Ranks[0] = r
			} else if val == 2 {
				ret.Type = hand.HandType_TwoPair
				if ret.Ranks[0] < r {
					ret.Ranks[1] = ret.Ranks[0]
					ret.Ranks[0] = r
				} else {
					ret.Ranks[1] = r
				}
			} else if val == 1 {
				if ret.Ranks[2] < r {
					ret.Ranks[3] = ret.Ranks[2]
					ret.Ranks[2] = r
				} else {
					ret.Ranks[3] = r
				}
			}
		}
		if ret.Type == hand.HandType_ThreeOfAKind {
			// Shifting ranks 2 to rank 1 and rand 3 to rank 2
			ret.Ranks[1] = ret.Ranks[2]
			ret.Ranks[2] = ret.Ranks[3]
			ret.Ranks[3] = card.CardRank_None
		}
		return ret, nil
	}

	if len(rankCount) == 4 {
		// This is a pair
		ret.Type = hand.HandType_Pair
		for r, val := range rankCount {
			if val == 2 {
				ret.Ranks[0] = r
			} else if val == 1 {
				if ret.Ranks[1] < r {
					ret.Ranks[3] = ret.Ranks[2]
					ret.Ranks[2] = ret.Ranks[1]
					ret.Ranks[1] = r
				} else if ret.Ranks[2] < r {
					ret.Ranks[3] = ret.Ranks[2]
					ret.Ranks[2] = r
				}
			}
		}
	}

	// Return single card type (nothing to do, ranks already calculated upper)
	return ret, nil
}

// GetBest gets the indexes of the best input hands. If there are several, they are equal.
// There must be at most 10 hands or ErrBadRequest shall be returned
func (d *Default) GetBest(ctx context.Context, hands hand.Hands) ([]int, error) {
	panic("not implemented") // TODO: Implement
}
