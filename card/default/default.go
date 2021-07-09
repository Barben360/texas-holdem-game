package dft

import (
	"context"
	"math/rand"
	"sort"

	"github.com/Barben360/texas-holdem-game/card"
	"github.com/Barben360/texas-holdem-game/errors"
)

// Default is the default carder implementation
type Default struct {
}

// New creates a new instance of default implementation of carder
func New(ctx context.Context) (card.Carder, error) {
	return new(Default), nil
}

// Shuffle shuffles cards
func (d *Default) Shuffle(ctx context.Context, cards card.Cards) error {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return nil
}

// Draw draws cards from top of pile. If not enough cards, draws all the cards and returns ErrResourceExhausted
func (d *Default) Draw(ctx context.Context, quantity int, cards card.Cards) (drawn card.Cards, remaining card.Cards, err error) {
	if quantity > len(cards) {
		return cards, card.Cards{}, errors.Errorf(errors.ErrResourceExhausted, "not enough cards to draw")
	}
	return cards[:quantity], cards[quantity:], nil
}

// Sort cards sorts cards by rank first then by suit
func (d *Default) Sort(ctx context.Context, cards card.Cards, descendent bool) error {
	if descendent {
		sort.Sort(card.CardsDesc(cards))
	} else {
		sort.Sort(card.CardsAsc(cards))
	}
	return nil
}

// Compare compares two cards by rank
// If equal is true, then ranks are the same
// If comp is false, then card1 < card2, otherwise card1 > card2. False if equal is true
func (d *Default) Compare(ctx context.Context, card1 card.Card, card2 card.Card) (equal bool, comp bool, err error) {
	if card1.Rank == card2.Rank {
		return true, false, nil
	}
	return false, card1.Rank < card2.Rank, nil
}
