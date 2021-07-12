package dft

import (
	"context"
	"math/rand"
	"sort"

	"github.com/Barben360/texas-holdem-game/card"
)

// Default is the default carder implementation
type Default struct {
}

// New creates a new instance of default implementation of carder
func New(ctx context.Context) (card.Carder, error) {
	return new(Default), nil
}

// Shuffle shuffles cards
func (d *Default) Shuffle52(ctx context.Context, cards *card.Cards52) error {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return nil
}

func (d *Default) Sort(ctx context.Context, cards card.Cards, descendent bool) error {
	if descendent {
		sort.Slice(cards, func(i, j int) bool {
			if cards[i].Rank == cards[j].Rank {
				return cards[i].Suit > cards[j].Suit
			}
			return cards[i].Rank > cards[j].Rank
		})
	} else {
		sort.Slice(cards, func(i, j int) bool {
			if cards[i].Rank == cards[j].Rank {
				return cards[i].Suit < cards[j].Suit
			}
			return cards[i].Rank < cards[j].Rank
		})
	}
	return nil
}
