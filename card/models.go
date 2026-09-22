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

// List of cards
type Cards []Card

// List of 2 cards
type Cards2 [2]Card

// List of 5 cards
type Cards5 [5]Card

// List of 7 cards
type Cards7 [7]Card

// List of 52 cards
type Cards52 [52]Card
