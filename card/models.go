package card

// CardSuit is the suit of a card
type CardSuit int8

const (
	CardSuit_Club    CardSuit = 0
	CardSuit_Diamond CardSuit = 1
	CardSuit_Hearts  CardSuit = 2
	CardSuit_Spades  CardSuit = 3
)

// CardRank is the rank of a card, ordered by value (as for poker)
type CardRank int8

const (
	CardRank_2     CardRank = 0
	CardRank_3     CardRank = 1
	CardRank_4     CardRank = 2
	CardRank_5     CardRank = 3
	CardRank_6     CardRank = 4
	CardRank_7     CardRank = 5
	CardRank_8     CardRank = 6
	CardRank_9     CardRank = 7
	CardRank_10    CardRank = 8
	CardRank_Jack  CardRank = 9
	CardRank_Queen CardRank = 10
	CardRank_King  CardRank = 11
	CardRank_Ace   CardRank = 12
)

// Card represents a card
type Card struct {
	Suit CardSuit
	Rank CardRank
}

// Represents a list of cards, top of the pile first
type Cards []Card
