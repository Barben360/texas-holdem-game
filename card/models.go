package card

// CardSuit is the suit of a card
type CardSuit int8

const (
	CardSuit_None    CardSuit = 0
	CardSuit_Club    CardSuit = 1
	CardSuit_Diamond CardSuit = 2
	CardSuit_Hearts  CardSuit = 3
	CardSuit_Spades  CardSuit = 4
)

// CardRank is the rank of a card, ordered by value (as for poker)
type CardRank int8

const (
	CardRank_None  CardRank = 0
	CardRank_2     CardRank = 1
	CardRank_3     CardRank = 2
	CardRank_4     CardRank = 3
	CardRank_5     CardRank = 4
	CardRank_6     CardRank = 5
	CardRank_7     CardRank = 6
	CardRank_8     CardRank = 7
	CardRank_9     CardRank = 8
	CardRank_10    CardRank = 9
	CardRank_Jack  CardRank = 10
	CardRank_Queen CardRank = 11
	CardRank_King  CardRank = 12
	CardRank_Ace   CardRank = 13
)

// Card represents a card
type Card struct {
	Suit CardSuit
	Rank CardRank
}

// Represents a list of cards, top of the pile first
type Cards2 [2]Card
type Cards5 [5]Card
type Cards7 [7]Card
type Cards52 [52]Card

type CardsAsc2 Cards2
type CardsDesc2 Cards2

type CardsAsc5 Cards5
type CardsDesc5 Cards5

func (c CardsAsc2) Len() int      { return len(c) }
func (c CardsAsc2) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsAsc2) Less(i, j int) bool {
	if c[i].Rank == c[j].Rank {
		return c[i].Suit < c[j].Suit
	}
	return c[i].Rank < c[j].Rank
}

func (c CardsDesc2) Len() int      { return len(c) }
func (c CardsDesc2) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsDesc2) Less(i, j int) bool {
	if c[i].Rank == c[j].Rank {
		return c[i].Suit > c[j].Suit
	}
	return c[i].Rank > c[j].Rank
}

func (c CardsAsc5) Len() int      { return len(c) }
func (c CardsAsc5) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsAsc5) Less(i, j int) bool {
	if c[i].Rank == c[j].Rank {
		return c[i].Suit < c[j].Suit
	}
	return c[i].Rank < c[j].Rank
}

func (c CardsDesc5) Len() int      { return len(c) }
func (c CardsDesc5) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsDesc5) Less(i, j int) bool {
	if c[i].Rank == c[j].Rank {
		return c[i].Suit > c[j].Suit
	}
	return c[i].Rank > c[j].Rank
}
