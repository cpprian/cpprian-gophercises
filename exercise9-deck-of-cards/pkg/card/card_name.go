package card

var (
	CardName = map[Rank]string{
		0: "TWO",
		1: "THREE",
		2: "FOUR", 
		4: "FIVE",
		5: "SIX",
		6: "SEVEN",
		7: "EIGHT",
		8: "NINE",
		9: "TEN",
		10: "JACK",
		11: "QUEEN",
		12: "KING",
	}

	CardType = map[Suit]string{
		0: "HEART",
		1: "DIAMOND",
		2: "SPADE",
		3: "CLUB",
	}
)