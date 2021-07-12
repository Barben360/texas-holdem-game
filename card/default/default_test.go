package dft

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Barben360/texas-holdem-game/card"
)

const nShuffles int = 1000

var carder *Default

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	c, err := New(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
	carder = c.(*Default)
	os.Exit(m.Run())
}
func TestShuffle52(t *testing.T) {
	var cards, cardsCopy card.Cards52
	k := 0
	for i := card.CardSuit_Club; i <= card.CardSuit_Spades; i++ {
		for j := card.CardRank_2; j <= card.CardRank_Ace; j++ {
			cards[k] = card.Card{
				Suit: i,
				Rank: j,
			}
			cardsCopy[k] = cards[k]
			k++
		}
	}

	// Shuffling
	for i := 0; i < nShuffles; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		err := carder.Shuffle52(ctx, &cards)
		cancel()
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(cards, cardsCopy) {
			t.Fatal("shuffling cards led to no entropy increase")
		}
		// Copying cards for next step
		for i, c := range cards {
			cardsCopy[i] = c
		}
	}
}

func TestSort(t *testing.T) {
	testData := []struct {
		InputCards          card.Cards
		Descendent          bool
		ExpectedSortedCards card.Cards
	}{
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
			},
			Descendent: false,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
			},
		},
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
			Descendent: false,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
			},
		},
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
			},
			Descendent: true,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
		},
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
			Descendent: true,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_King,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
		},

		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
			},
			Descendent: false,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
			},
		},
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
			Descendent: false,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
			},
		},
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
			},
			Descendent: true,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
		},
		{
			InputCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
			Descendent: true,
			ExpectedSortedCards: card.Cards{
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Diamond,
				},
				{
					Rank: card.CardRank_2,
					Suit: card.CardSuit_Club,
				},
			},
		},
	}

	for _, td := range testData {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		err := carder.Sort(ctx, td.InputCards, td.Descendent)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(td.InputCards, td.ExpectedSortedCards) {
			t.Fatalf("expected input cards and sorted cards are not the same\n+ %+v\n- %+v", td.ExpectedSortedCards, td.InputCards)
		}
	}
}
