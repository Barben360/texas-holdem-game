package dft

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Barben360/texas-holdem-game/card"
	carder_dft "github.com/Barben360/texas-holdem-game/card/default"
	"github.com/Barben360/texas-holdem-game/hand"
	"github.com/Barben360/texas-holdem-game/services"
)

var hander *Default

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	svcs, err := services.Init(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	carder, err := carder_dft.New(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	h, err := New(ctx, svcs, carder)
	cancel()
	if err != nil {
		os.Exit(1)
	}
	hander = h.(*Default)
	os.Exit(m.Run())
}

func TestCheckBase(t *testing.T) {
	testData := []struct {
		InputCards   card.Cards5
		ExpectedHand hand.Hand
	}{
		// Single card
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Club,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_Jack,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_6,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_3,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_SingleCard,
				Suit: card.CardSuit_None,
				Ranks: [5]card.CardRank{
					card.CardRank_Jack,
					card.CardRank_6,
					card.CardRank_5,
					card.CardRank_3,
					card.CardRank_2,
				},
			},
		},
		// Two pair
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Club,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_6,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_5,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_TwoPair,
				Suit: card.CardSuit_None,
				Ranks: [5]card.CardRank{
					card.CardRank_5,
					card.CardRank_2,
					card.CardRank_6,
					card.CardRank_None,
					card.CardRank_None,
				},
			},
		},
		// Three-of-a-kind
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Club,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_Jack,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_ThreeOfAKind,
				Suit: card.CardSuit_None,
				Ranks: [5]card.CardRank{
					card.CardRank_2,
					card.CardRank_Jack,
					card.CardRank_5,
					card.CardRank_None,
					card.CardRank_None,
				},
			},
		},
		// Straight
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Club,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_6,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_4,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_3,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_Straight,
				Suit: card.CardSuit_None,
				Ranks: [5]card.CardRank{
					card.CardRank_6,
					card.CardRank_5,
					card.CardRank_4,
					card.CardRank_3,
					card.CardRank_2,
				},
			},
		},
		// Flush
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_Jack,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_6,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_Ace,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_3,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_Flush,
				Suit: card.CardSuit_Diamond,
				Ranks: [5]card.CardRank{
					card.CardRank_Ace,
					card.CardRank_Jack,
					card.CardRank_6,
					card.CardRank_5,
					card.CardRank_3,
				},
			},
		},
		// Full house
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Club,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_2,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_5,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_FullHouse,
				Suit: card.CardSuit_None,
				Ranks: [5]card.CardRank{
					card.CardRank_2,
					card.CardRank_5,
					card.CardRank_None,
					card.CardRank_None,
					card.CardRank_None,
				},
			},
		},
		// Four-of-a-kind
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Club,
					Rank: card.CardRank_Ace,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_Ace,
				},
				{
					Suit: card.CardSuit_Diamond,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_Ace,
				},
				{
					Suit: card.CardSuit_Spades,
					Rank: card.CardRank_Ace,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_FourOfAKind,
				Suit: card.CardSuit_None,
				Ranks: [5]card.CardRank{
					card.CardRank_Ace,
					card.CardRank_5,
					card.CardRank_None,
					card.CardRank_None,
					card.CardRank_None,
				},
			},
		},
		// Straight flush
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Spades,
					Rank: card.CardRank_5,
				},
				{
					Suit: card.CardSuit_Spades,
					Rank: card.CardRank_9,
				},
				{
					Suit: card.CardSuit_Spades,
					Rank: card.CardRank_6,
				},
				{
					Suit: card.CardSuit_Spades,
					Rank: card.CardRank_8,
				},
				{
					Suit: card.CardSuit_Spades,
					Rank: card.CardRank_7,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_StraightFlush,
				Suit: card.CardSuit_Spades,
				Ranks: [5]card.CardRank{
					card.CardRank_9,
					card.CardRank_8,
					card.CardRank_7,
					card.CardRank_6,
					card.CardRank_5,
				},
			},
		},
		// Royal flush
		{
			InputCards: card.Cards5{
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_Ace,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_Jack,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_King,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_Queen,
				},
				{
					Suit: card.CardSuit_Hearts,
					Rank: card.CardRank_10,
				},
			},
			ExpectedHand: hand.Hand{
				Type: hand.HandType_RoyalFlush,
				Suit: card.CardSuit_Hearts,
				Ranks: [5]card.CardRank{
					card.CardRank_Ace,
					card.CardRank_King,
					card.CardRank_Queen,
					card.CardRank_Jack,
					card.CardRank_10,
				},
			},
		},
	}

	for _, td := range testData {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		h, err := hander.check(ctx, td.InputCards)
		cancel()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(h, td.ExpectedHand) {
			t.Fatalf("hand is not as expected\n+ %+v\n- %+v", td.ExpectedHand, h)
		}
	}
}
