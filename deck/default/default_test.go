package dft

import (
	"context"
	"testing"
	"time"

	cardDft "github.com/Barben360/texas-holdem-game/card/default"
	"github.com/Barben360/texas-holdem-game/deck"
)

func TestDeck(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	carder, err := cardDft.New(ctx)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	decker, err := New(ctx, carder)
	cancel()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	_, err = decker.New(ctx, deck.Type_UNKNOWN)
	cancel()
	if err == nil {
		t.Fatal("creating new deck with unknown type should lead to an error")
	}
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	dck, err := decker.New(ctx, deck.Type_52)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	if len(dck.Cards) != 52 {
		t.Fatal("there should be 52 cards in deck")
	}

	// Testing shuffle
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	err = decker.Shuffle(ctx, dck)
	cancel()
	if err != nil {
		t.Fatal(err)
	}
	// We don't check result, it is already done by card shuffle function which is forwarded
}
