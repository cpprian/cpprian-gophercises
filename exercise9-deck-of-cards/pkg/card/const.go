package card

type Rank uint8
type Suit uint8

const (
	TWO Rank = iota
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	J
	Q
	K
	A
)

const (
	HEART Suit = iota
	DIAMOND
	SPADE
	CLUB
)

const (
	MINRANK Rank = 0
	MAXRANK Rank = 4
	MINSUIT Suit = 0
	MAXSUIT Suit = 13
)