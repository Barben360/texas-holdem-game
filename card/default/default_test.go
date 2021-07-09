package dft

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/errors"
)

func TestCard(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	carder, err := New(ctx)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	cards := card.Cards{
		{
			Suit: card.CardSuit_Club,
			Rank: card.CardRank_10,
		},
		{
			Suit: card.CardSuit_Spades,
			Rank: card.CardRank_Jack,
		},
		{
			Suit: card.CardSuit_Hearts,
			Rank: card.CardRank_2,
		},
		{
			Suit: card.CardSuit_Diamond,
			Rank: card.CardRank_4,
		},
		{
			Suit: card.CardSuit_Club,
			Rank: card.CardRank_Jack,
		},
		{
			Suit: card.CardSuit_Spades,
			Rank: card.CardRank_King,
		},
		{
			Suit: card.CardSuit_Club,
			Rank: card.CardRank_6,
		},
		{
			Suit: card.CardSuit_Spades,
			Rank: card.CardRank_5,
		},
	}

	cardsCopy := make(card.Cards, len(cards))
	copy(cardsCopy, cards)

	expectedSortedCards := card.Cards{
		{
			Suit: card.CardSuit_Hearts,
			Rank: card.CardRank_2,
		},
		{
			Suit: card.CardSuit_Diamond,
			Rank: card.CardRank_4,
		},
		{
			Suit: card.CardSuit_Spades,
			Rank: card.CardRank_5,
		},
		{
			Suit: card.CardSuit_Club,
			Rank: card.CardRank_6,
		},
		{
			Suit: card.CardSuit_Club,
			Rank: card.CardRank_10,
		},
		{
			Suit: card.CardSuit_Club,
			Rank: card.CardRank_Jack,
		},
		{
			Suit: card.CardSuit_Spades,
			Rank: card.CardRank_Jack,
		},
		{
			Suit: card.CardSuit_Spades,
			Rank: card.CardRank_King,
		},
	}

	// Testing shuffle. Order must change at least once (that would be very unlikely to happen several times)
	atLeastOneChange := false
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		err = carder.Shuffle(ctx, cards)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(cards, cardsCopy) {
			atLeastOneChange = true
			break
		}
	}
	if !atLeastOneChange {
		t.Fatal("shuffling cards failed (order has not changed)")
	}

	// Testing sort
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	err = carder.Sort(ctx, cards, true)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	if !func() bool {
		for i, c := range cards {
			if c != expectedSortedCards[len(expectedSortedCards)-i-1] {
				return false
			}
		}
		return true
	}() {
		t.Fatal("sorting by descending order failed")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	err = carder.Sort(ctx, cards, false)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(cards, expectedSortedCards) {
		t.Fatal("sorting by ascending order failed")
	}

	// Testing draw
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	drawn, remaining, err := carder.Draw(ctx, 3, cards)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	if len(drawn) != 3 || len(remaining) != 5 {
		t.Fatal("drawing 3 cards failed (inconsistent number of drawn or remaining cards)")
	}

	if !reflect.DeepEqual(drawn, cards[:3]) {
		t.Fatal("drawing 3 cards failed (inconsistent values for drawn cards)")
	}
	if !reflect.DeepEqual(remaining, cards[3:]) {
		t.Fatal("drawing 3 cards failed (inconsistent values for remaining cards)")
	}
	cards = remaining
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	drawn, remaining, err = carder.Draw(ctx, 10, cards)
	cancel()
	if err == nil {
		t.Fatal("there should be an error when drawing too many cards")
	}
	if e, ok := err.(*errors.GenericError); ok {
		if e.ErrType != errors.ErrResourceExhausted {
			t.Fatal("there should be an ErrResourceExhausted error when drawing too many cards")
		}
	} else {
		t.Fatal("error should be a GenericError")
	}
	if len(drawn) != 5 {
		t.Fatalf("there should be an 5 drawn cards, got %d", len(drawn))
	}
	if len(remaining) != 0 {
		t.Fatal("there should be no remaining cards")
	}

	// Testing compare
	compareTests := []struct {
		cardA, cardB                card.Card
		expectedEqual, expectedComp bool
	}{
		{
			cardA: card.Card{
				Suit: card.CardSuit_Club,
				Rank: card.CardRank_2,
			},
			cardB: card.Card{
				Suit: card.CardSuit_Club,
				Rank: card.CardRank_2,
			},
			expectedEqual: true,
			expectedComp:  false,
		},
		{
			cardA: card.Card{
				Suit: card.CardSuit_Club,
				Rank: card.CardRank_2,
			},
			cardB: card.Card{
				Suit: card.CardSuit_Diamond,
				Rank: card.CardRank_5,
			},
			expectedEqual: false,
			expectedComp:  true,
		},
		{
			cardA: card.Card{
				Suit: card.CardSuit_Diamond,
				Rank: card.CardRank_5,
			},
			cardB: card.Card{
				Suit: card.CardSuit_Club,
				Rank: card.CardRank_2,
			},
			expectedEqual: false,
			expectedComp:  false,
		},
	}
	for _, compareTest := range compareTests {
		ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
		equal, comp, err := carder.Less(ctx, compareTest.cardA, compareTest.cardB)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
		if equal != compareTest.expectedEqual {
			t.Fatalf("card compare equal value not as expected - %+v", compareTest)
		}
		if comp != compareTest.expectedComp {
			t.Fatalf("card compare comp value not as expected - %+v", compareTest)
		}
	}
}
